/**
 * Created by lfs
 * Email: 1144828910@qq.com
 * Date: 2021/3/16
 * Time: 9:18
 */

import {_handbook, getAesEncryptionMode,decryptionHandbookValue,CryptoMode,JJEncodeEval} from "./CipherText";

/**
 * @description 解析服务器返回的密文数据
 * @param response
 */
export default (response) => {
  let xKey           = CryptoMode.enc.Utf8.parse(response[decryptionHandbookValue(_handbook[8])].substr(0, 32))//aes key
  let xIv            = CryptoMode.enc.Utf8.parse(response[decryptionHandbookValue(_handbook[8])].substr(32, 16))//aes iv
  let encryptionData = CryptoMode.enc.Hex.parse(response[decryptionHandbookValue(_handbook[8])].substr(48)).toString(CryptoMode.enc.Base64)//密文数据转 base64
  let decryptData         = CryptoMode.AES.decrypt(encryptionData, xKey, getAesEncryptionMode(xIv))//解密的密文数据
  response[decryptionHandbookValue(_handbook[8])]      = JJEncodeEval(decryptionHandbookValue(_handbook[9])).parse(decryptData.toString(CryptoMode.enc.Utf8))
}
