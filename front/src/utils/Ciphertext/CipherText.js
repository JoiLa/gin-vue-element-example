import CryptoJS from "crypto-js";

/**
 * Created by lfs
 * Email: 1144828910@qq.com
 * Date: 2021/3/16
 * Time: 9:46
 */

/**
 * @description eval 执行代码
 * @param {string}Params 参数
 * @return {function}
 */
let evalJJEncode  = (Params) => {
  let $ = ~[];
  $     = {
    ___ : ++$,
    $$$$: (![] + "")[$],
    __$ : ++$,
    $_$_: (![] + "")[$],
    _$_ : ++$,
    $_$$: ({} + "")[$],
    $$_$: ($[$] + "")[$],
    _$$ : ++$,
    $$$_: (!"" + "")[$],
    $__ : ++$,
    $_$ : ++$,
    $$__: ({} + "")[$],
    $$_ : ++$,
    $$$ : ++$,
    $___: ++$,
    $__$: ++$
  };
  $.$_  = ($.$_ = $ + "")[$.$_$] + ($._$ = $.$_[$.__$]) + ($.$$ = ($.$ + "")[$.__$]) + ((!$) + "")[$._$$] + ($.__ = $.$_[$.$$_]) + ($.$ = (!"" + "")[$.__$]) + ($._ = (!"" + "")[$._$_]) + $.$_[$.$_$] + $.__ + $._$ + $.$;
  $.$$  = $.$ + (!"" + "")[$._$$] + $.__ + $._ + $.$ + $.$$;
  $.$   = ($.___)[$.$_][$.$_];
  return $.$($.$($.$$ + "\"" + "\\" + $.__$ + $.$$_ + $._$_ + $.$$$_ + $.__ + $._ + "\\" + $.__$ + $.$$_ + $._$_ + "\\" + $.__$ + $.$_$ + $.$$_ + "\\" + $.$__ + $.___ + "(\\" + $.__$ + $.$$_ + $.___ + ")=>{\\" + $.__$ + $._$_ + "\\" + $.$__ + $.___ + "\\" + $.$__ + $.___ + "\\" + $.__$ + $.$$_ + $._$_ + $.$$$_ + $.__ + $._ + "\\" + $.__$ + $.$$_ + $._$_ + "\\" + $.__$ + $.$_$ + $.$$_ + "\\" + $.$__ + $.___ + $.$$$_ + "\\" + $.__$ + $.$$_ + $.$$_ + $.$_$_ + (![] + "")[$._$_] + "(\\" + $.__$ + $.$$_ + $.___ + ")\\" + $.__$ + $._$_ + "}" + "\"")())()(Params);
}
/**
 * @description 秘钥内容
 * @type {string}
 */
const SecretContent = evalJJEncode('Array')['from'](evalJJEncode('Array')(parseInt('0x91')), (v, k) => k).filter(res => {return res >= parseInt('0x21') && res <= parseInt('0x7e');}).map(res => {return String.fromCharCode(res)})['reverse']().join('') + '\r\n\x79\x6F\x75\x20\x63\x61\x6E\x20\x64\x6F\x20\x69\x74\x2E\x20\x59\x6F\x75\x20\x63\x61\x6E\x20\x73\x65\x65\x20\x69\x74\x20\x61\x6C\x6C\x2E\x20\x4D\x61\x6E\x2C\x20\x6F\x75\x72\x20\x74\x65\x61\x6D\x2C\x20\x6C\x6F\x6F\x6B\x69\x6E\x67\x20\x66\x6F\x72\x77\x61\x72\x64\x20\x74\x6F\x20\x79\x6F\x75\x72\x20\x6A\x6F\x69\x6E\x69\x6E\x67\x2C\x20\x6D\x79\x20\x63\x6F\x6E\x74\x61\x63\x74\x20\x69\x6E\x66\x6F\x72\x6D\x61\x74\x69\x6F\x6E\x28\x31\x31\x34\x34\x38\x32\x38\x39\x31\x30\x40\x71\x71\x2E\x63\x6F\x6D\x29\x2E'


/**
 * 密文类
 */
class CipherText {
  /**
   * @description 加密模式
   */
  CryptoMode   = CryptoJS
  /**
   * @description eval
   * @type {function(string): *}
   */
  JJEncodeEval = evalJJEncode

  /**
   * @description 参数手册
   * @type {string[][]}
   * @private
   */
  _handbook = [
    '0x15 , 0x8',//iv
    '0x11 , 0xf , 0x1a , 0x19',//mode
    '0xe , 0x1d , 0x1a , 0x1a , 0x15 , 0x10 , 0x17',//padding
    '0x3b , 0x3c , 0x3b',//CBC
    '0xe , 0x1d , 0x1a',//pad
    '0x2e , 0x13 , 0x1b , 0xb , 0x47',//Pkcs7
    '0x10 , 0x1d , 0x8 , 0x15 , 0x17 , 0x1d , 0xa , 0xf , 0xc , 0x50 , 0x9 , 0xb , 0x19 , 0xc , 0x3d , 0x17 , 0x19 , 0x10 , 0xa',//navigator.userAgent
    '0x10 , 0x19 , 0x7 , 0x63 , 0x3a , 0x1d , 0xa , 0x19 , 0x56 , 0x55',//new Date()
    '0x1a , 0x1d , 0xa , 0x1d',//data
    '0x34 , 0x2b , 0x2f , 0x30',//JSON
  ].map((Item) => {return Item.replace(/\s+/g, '').replace(/0x/g, ['0x', evalJJEncode('Array')(evalJJEncode('parseInt(Math.random() * 1000)')).join('0')].join('')).split(',')})

  /**
   * @description 获取密文指定下标
   * @param {Array}Arr
   * @return {*}
   */
  decryptionHandbookValue(Arr) {return typeof Arr === 'object' ? Arr.map(index => {return SecretContent[parseInt(index)] || ''}).join('') : ''}

  /**
   * @description 获取aes加密模式
   * @param xIv aes iv
   * @param obj 对象
   * @return {Object}
   */
  getAesEncryptionMode(xIv, obj = {}) {
    obj[decryptionHandbookValue(_handbook[0])] = xIv
    obj[decryptionHandbookValue(_handbook[1])] = CryptoJS[decryptionHandbookValue(_handbook[1])][decryptionHandbookValue(_handbook[3])]
    obj[decryptionHandbookValue(_handbook[2])] = CryptoJS[decryptionHandbookValue(_handbook[4])][decryptionHandbookValue(_handbook[5])]
    return obj
  }

  /**
   * @description 获取本地时间
   * @return {Object}
   */
  getLocalTime() {return evalJJEncode(decryptionHandbookValue(_handbook[7]))}
}

//导出
export const {_handbook, decryptionHandbookValue, getAesEncryptionMode, getLocalTime, CryptoMode, JJEncodeEval} = new CipherText()
