/**
 * description:
 * author: lfs <1144828910@qq.com>
 * Date: 2021/3/10
 */
package globals

import (
    `crypto/aes`
    `crypto/cipher`
    `crypto/md5`
    `encoding/hex`
    `fmt`
    `net/url`
    `sort`
    `strings`

    `Api/pkg/e`
    `Api/pkg/helper`
    `github.com/gin-gonic/gin`
    `github.com/manucorporat/try`
)

/**
 * @Description: 解析密文
 * @return gin.HandlerFunc
 */
func ParseCiphertextMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        try.This(func() {
            data := c.PostForm("data")
            if data == "" {
                panic("send data failed")
            }
            checkRequestSignIsNormal(c, data)
            decryptData := aesDecryptCBC(data)
            fmt.Println("用户请求的数据->", decryptData)
            parseData, err := url.ParseQuery(decryptData)
            if err != nil {
                panic("submit data parse failed")
            }
            c.Request.PostForm = parseData // 将数据保存到 PostForm
            c.Next()
        }).Catch(func(err try.E) {
            // 反馈数据
            helper.NewResult(c).Error(e.InvalidParams, e.GetMsg(e.InvalidParams))
        })

    }
}

/**
 * @Description: 检查请求的签名是否正确
 * @param Client 客户端
 * @param Ciphertext 密文
 * @return error
 */
func checkRequestSignIsNormal(Client *gin.Context, Ciphertext string) {
    v, err := url.ParseQuery(Client.Request.URL.RawQuery)
    if err != nil {
        panic(err)
    }
    var queryKeyArr string
    var queryValueArr string
    for i := range v {
        queryKeyArr += i
        queryValueArr += v.Get(i)
    }
    strArr := []string{Client.Request.URL.Path, Client.Request.Method, Client.Request.Header.Get("User-Agent"), queryKeyArr, queryValueArr}
    sort.Strings(strArr) // 升序排序
    sign := fmt.Sprintf("%x", md5.Sum([]byte(strings.ToUpper(strings.Join(strArr, ",")))))
    if sign != Ciphertext[len(Ciphertext)-32:] {
        // 签名不一致
        panic("sign is failed")
    }
}

/**
 * @Description: aes解密
 * @param CryptoData
 * @return string
 */
func aesDecryptCBC(CryptoData string) string {
    if len(CryptoData) < 96 {
        panic("crypto data defined failed")
    }
    kv := CryptoData[:96] // 从提交的报文中，获取出 aes秘钥及 iv
    if len(kv) != 96 {
        panic("kev iv is defined failed")
    }
    key, err := decodeHexUpper(kv[:64])
    if err != nil {
        panic(err)
    }
    iv, err := decodeHexUpper(kv[64:])
    if err != nil {
        panic(err)
    }
    EncryptData := CryptoData[96:] // 从提交的报文中，获取出加密的数据
    if len(EncryptData) < 32 {
        panic("encrypt data defined failed")
    }
    // 客户端传过来的aes hex 密文数据
    cryptoByte, err := decodeHexUpper(EncryptData[:len(EncryptData)-32])
    if err != nil {
        panic(err)
    }
    block, err := aes.NewCipher(key)
    if err != nil {
        panic("aes failed")
    }
    blockMode := cipher.NewCBCDecrypter(block, iv) // 加密模式
    origData := make([]byte, len(cryptoByte))      // 创建数组
    blockMode.CryptBlocks(origData, cryptoByte)    // 解密
    origData = pKCS5UnPadding(origData)
    // ecb mode
    /*blockMode := newECBDecrypt(block)
      origData := make([]byte, len(cryptoByte))
      blockMode.CryptBlocks(origData, cryptoByte)
      origData = pKCS5UnPadding(origData)*/
    return string(origData)
}

type ecb struct {
    b         cipher.Block
    blockSize int
}

func newECB(b cipher.Block) *ecb {
    return &ecb{
        b:         b,
        blockSize: b.BlockSize(),
    }
}

type ecbDecrypt ecb

// newECBDecrypt returns a BlockMode which decrypts in electronic code book
// mode, using the given Block.
func newECBDecrypt(b cipher.Block) cipher.BlockMode {
    return (*ecbDecrypt)(newECB(b))
}
func (x *ecbDecrypt) BlockSize() int { return x.blockSize }
func (x *ecbDecrypt) CryptBlocks(dst, src []byte) {
    if len(src)%x.blockSize != 0 {
        panic("crypto/cipher: input not full blocks")
    }
    if len(dst) < len(src) {
        panic("crypto/cipher: output smaller than input")
    }
    for len(src) > 0 {
        x.b.Decrypt(dst, src[:x.blockSize])
        src = src[x.blockSize:]
        dst = dst[x.blockSize:]
    }
}

/**
 * @Description: pkcs5填充
 * @param origData
 * @return []byte
 */
func pKCS5UnPadding(origData []byte) []byte {
    length := len(origData)
    unPadding := int(origData[length-1])
    return origData[:(length - unPadding)]
}

// 16进制字符串转换成byte
func decodeHexUpper(str string) ([]byte, error) {
    return hex.DecodeString(strings.ToLower(str))
}
