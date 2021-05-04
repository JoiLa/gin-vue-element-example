/**
 * description:
 * author: lfs <1144828910@qq.com>
 * Date: 2021/3/15
 */
package globals

import (
    `bytes`
    `crypto/aes`
    `crypto/cipher`
    `crypto/md5`
    `encoding/hex`
    `encoding/json`
    `fmt`
    `net/http`
    `time`

    `Api/pkg/e`
    `github.com/gin-gonic/gin`
)

/**
 * @Description: 生成密文
 * @return gin.HandlerFunc
 */
func GenerateCiphertextMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 请求前
        response, exists := c.Get(e.ParamsResponseJson)
        if !exists {
            c.Next() // 执行控制台方法
            response, exists = c.Get(e.ParamsResponseJson)
        }
        // 判断是否存在
        if exists {
            // 判断响应类型
            switch response.(type) {
            case gin.H:
                // 响应的是 数据结构体
                // 将数据结构体转 字符
                byteData, err := json.Marshal(response)
                if err != nil {
                    c.Status(http.StatusInternalServerError) // service errors
                } else {
                    key := randomGenerateAesKey()
                    iv := randomGenerateAesIv()
                    encryptResult := []byte(encodeHexUpper(aesEncryptCBC(byteData, key, iv)))
                    var buffer bytes.Buffer
                    buffer.Write(key)
                    buffer.Write(iv)
                    buffer.Write(encryptResult)
                    c.Data(http.StatusOK, "application/octet-stream", buffer.Bytes())
                    c.Abort()
                }
                break
            }
        }
    }
}

/**
 * @Description: 随机生成aes Key
 * @return []byte
 */
func randomGenerateAesKey() []byte {
    // 长度必须为16, 24或者32
    return []byte(fmt.Sprintf("%x", md5.Sum([]byte(time.Now().String()+"join this data,full aes key"))))
}

/**
 * @Description: 随机生成aes iv
 * @return []byte
 */
func randomGenerateAesIv() []byte {
    // 长度必须为16, 24或者32
    iv := fmt.Sprintf("%x", md5.Sum([]byte(time.Now().String()+"join this data,full aes iv")))
    return []byte(iv[:16])
}

/**
 * @Description: aesEncryptCBC aes cbc模式加密
 * @param OrigData
 * @param Key
 * @param Iv
 * @return EncryptedResult
 */
func aesEncryptCBC(OrigData, Key, Iv []byte) (EncryptedResult []byte) {
    // 分组秘钥
    // NewCipher该函数限制了输入k的长度必须为16, 24或者32
    block, _ := aes.NewCipher(Key)
    blockSize := block.BlockSize()                   // 获取秘钥块的长度
    OrigData = pkcs5Padding(OrigData, blockSize)     // 补全码
    blockMode := cipher.NewCBCEncrypter(block, Iv)   // 加密模式
    EncryptedResult = make([]byte, len(OrigData))    // 创建数组
    blockMode.CryptBlocks(EncryptedResult, OrigData) // 加密
    return EncryptedResult
}
func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
    padding := blockSize - len(ciphertext)%blockSize
    padText := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(ciphertext, padText...)
}

// byte转换16进制字符串
func encodeHexUpper(byte []byte) string {
    return hex.EncodeToString(byte)
}
