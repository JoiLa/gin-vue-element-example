/**
 * Created by lfs
 * Email: 1144828910@qq.com
 * Date: 2021/3/12
 * Time: 22:03
 */
import qs                                                                                   from 'qs'
import {_handbook, decryptionHandbookValue, getAesEncryptionMode, getLocalTime, CryptoMode,JJEncodeEval} from './CipherText'


/**
 * @description aes秘钥偏移
 * @type {string}
 */
const aesKeyOffset  = CryptoMode.MD5([getLocalTime().getTime().toString(), _handbook.join('')]['reverse']().join('')).toString()
/**
 * @description 设置空数据的默认值
 * @param Prop 参数
 * @param DefaultValue 默认数据
 * @return {*}
 */
const setEmptyValue = (Prop, DefaultValue) => {return Prop || DefaultValue}
/**
 * @description 加密客户端发送的数据
 * @param config
 * @return {*}
 */
export default (config) => {
  let timestamp           = (getLocalTime()).getTime().toString()//aes的key基值
  config.params           = setEmptyValue(config.params, {})
  config.data             = setEmptyValue(config.data, {})
  config.params.timestamp = setEmptyValue(config.params.timestamp, timestamp)
  config.data.timestamp   = setEmptyValue(config.data.timestamp, timestamp)

  let sourceData = CryptoMode.enc.Utf8.parse(qs.stringify(config.data))//发送的源POST数据
  let xKey       = CryptoMode.enc.Utf8.parse(CryptoMode.MD5([timestamp, aesKeyOffset].join('')).toString())//秘钥。长度32的16进制字符串。
  let xIv        = CryptoMode.lib.WordArray.random(128 / 8)//随机生成长度为32的16进制字符串。IV称为初始向量，不同的IV加密后的字符串是不同的，加密和解密需要相同的IV。
  let encrypted  = CryptoMode.AES.encrypt(sourceData, xKey, getAesEncryptionMode(xIv))//加密后的数据
  //计算提交的数据签名
  let sign                                        = CryptoMode.MD5(
    [
      config.url,
      config.method,
      JJEncodeEval(decryptionHandbookValue(_handbook[6])),
      Object.keys(config.params).join(''),
      Object.values(config.params).join('')
    ].sort().join(',').toUpperCase()
  ).toString()
  //设置提交的数据
  let sendData                                    = {}
  //封装发送的数据
  sendData[decryptionHandbookValue(_handbook[8])] = [
    [
      xKey.toString(),
      xIv.toString(CryptoMode.enc.Hex)
    ].join(''),
    encrypted.ciphertext.toString(),
    sign
  ].join('')
  //设置发送的数据到 axios中
  config[decryptionHandbookValue(_handbook[8])]   = qs.stringify(sendData)
  //返回
  return config
}
