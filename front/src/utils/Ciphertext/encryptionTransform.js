/**
 * Created by lfs
 * Email: 1144828910@qq.com
 * Date: 2021/3/10
 * Time: 13:46
 */

let Secret           = Array.from(Array(0x91), (v, k) => k).filter(res => {return res >= 0x21 && res <= 0x7e;}).map(res => {return String.fromCharCode(res)}).reverse().join('') + "\r\n\x79\x6F\x75\x20\x63\x61\x6E\x20\x64\x6F\x20\x69\x74\x2E\x20\x59\x6F\x75\x20\x63\x61\x6E\x20\x73\x65\x65\x20\x69\x74\x20\x61\x6C\x6C\x2E\x20\x4D\x61\x6E\x2C\x20\x6F\x75\x72\x20\x74\x65\x61\x6D\x2C\x20\x6C\x6F\x6F\x6B\x69\x6E\x67\x20\x66\x6F\x72\x77\x61\x72\x64\x20\x74\x6F\x20\x79\x6F\x75\x72\x20\x6A\x6F\x69\x6E\x69\x6E\x67\x2C\x20\x6D\x79\x20\x63\x6F\x6E\x74\x61\x63\x74\x20\x69\x6E\x66\x6F\x72\x6D\x61\x74\x69\x6F\x6E\x28\x31\x31\x34\x34\x38\x32\x38\x39\x31\x30\x40\x71\x71\x2E\x63\x6F\x6D\x29\x2E";
//获取密文指定下标
const getSecretValue = (t) => {return t.map(res => {return Secret[res] || ''}).join('')}

/**
 *
 * @param {string} Params
 * @returns {string}
 */
const getCipherText = (Params) => {
  let arr = Params.split('')
  if (arr.length <= 0) return []
  return arr.map((Item) => {
    let result = Secret.indexOf(Item)
    if (result === -1) {
      //未找到索引
      return null
    } else {
      //找到索引，返回hex
      return ['0x', Secret.indexOf(Item).toString(16)].join('')
    }
  }).filter(i => i).join(' , ')
}
let encryptText            = getCipherText("JSON")
console.log(encryptText)
/*let decryptText            = getSecretValue([0x52, 0x4f, 0x4b, 0x48, 0x53])
console.log(decryptText)*/
