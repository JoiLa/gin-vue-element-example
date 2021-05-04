/**
 * Created by lfs
 * Email: 1144828910@qq.com
 * Date: 2021/3/15
 * Time: 8:32
 */
const processStartTime = new Date() // 处理开始时间
const JavaScriptObfuscator = require('javascript-obfuscator')// 引入->混淆第三方包
const fs = require('fs')// 引入->文件操作
const slog = require('single-line-log').stdout // 日志
const path = require('path')
const yargs = require('yargs')
// argv 对象用来保存命令行参数，传递参数时，参数名以 -- 开头，中间使用 = 或 空格，然后接上值 。
let args = yargs
  .options(
    {
      dir: {
        describe: '指定混淆的目录名（当前目录下的文件名）',
        demandOption: true,
        type: 'string'
      }
    }
  ).argv
// 处理的路径
const processPath = path.join(process.cwd(), args.dir, path.sep)
// 判断目录文件是否存在
if (!fs.existsSync(processPath) || fs.lstatSync(processPath).isFile(processPath)) {
  console.log('\x1B[31m%s\x1B[0m', `\r\n地址->${processPath}\t提示：不是目录或者目录不存在`)
  // 目录不存在
  process.exit()
}

// 封装的 进度条 工具
function ProgressBar(description, bar_length) {
  // 两个基本参数(属性)
  this.description = description || 'Progress' // 命令行开头的文字信息
  this.length = bar_length || 25 // 进度条的长度(单位：字符)，默认设为 25

  // 刷新进度条图案、文字的方法
  this.render = (opts) => {
    if (opts.total <= 0) return
    const percent = (opts.completed / opts.total).toFixed(4) // 计算进度(子任务的 完成数 除以 总数)
    const cell_num = Math.floor(percent * this.length) // 计算需要多少个 █ 符号来拼凑图案

    // 拼接黑色条
    let cell = ''
    for (let i = 0; i < cell_num; i++) cell += '█'

    // 拼接灰色条
    let empty = ''
    for (let i = 0; i < this.length - cell_num; i++) empty += '░'

    // 拼接最终文本
    const cmdText = `${this.description}: ${(100 * percent).toFixed(2)}% ${cell}${empty} ${opts.completed}/${opts.total} ${opts.fileName ? `uploading : ${opts.fileName}` : ''} ${opts.fileDir ? 'complete: ' + opts.fileDir : ''}`
    // 在单行输出文本
    slog(cmdText)
  }
}

/**
 * @description 加密配置
 * @type {{middle: {}, max: {}, lower: {}}}
 */
const obfuscationEncryptionConfig = {
  lower: {// 性能会比没有混淆的情况稍微慢一些
    compact: true,
    controlFlowFlattening: false,
    deadCodeInjection: false,
    debugProtection: false,
    debugProtectionInterval: false,
    disableConsoleOutput: false,
    identifierNamesGenerator: 'hexadecimal',
    log: false,
    numbersToExpressions: false,
    renameGlobals: false,
    rotateStringArray: true,
    selfDefending: false,
    shuffleStringArray: true,
    simplify: true,
    splitStrings: false,
    stringArray: true,
    stringArrayEncoding: ['base64'],
    stringArrayIndexShift: true,
    stringArrayWrappersCount: 1,
    stringArrayWrappersChainedCalls: true,
    stringArrayWrappersParametersMaxCount: 2,
    stringArrayWrappersType: 'variable',
    stringArrayThreshold: 0.75,
    unicodeEscapeSequence: false
  },
  middle: {// 性能将比没有混淆的情况下降低30-35%
    compact: true,
    controlFlowFlattening: true,
    controlFlowFlatteningThreshold: 0.5,
    deadCodeInjection: true,
    deadCodeInjectionThreshold: 0.4,
    debugProtection: false,
    debugProtectionInterval: false,
    disableConsoleOutput: true,
    identifierNamesGenerator: 'hexadecimal',
    log: false,
    renameGlobals: false,
    rotateStringArray: true,
    selfDefending: true,
    stringArray: true,
    stringArrayEncoding: ['base64'],
    stringArrayThreshold: 0.75,
    transformObjectKeys: true,
    unicodeEscapeSequence: false
  },
  max: {// 性能将比没有混淆的情况下慢50-100%
    compact: true, // 压缩代码
    controlFlowFlattening: true, // 是否启用控制流扁平化(降低1.5倍的运行速度)
    controlFlowFlatteningThreshold: 0.1, // 应用概率;在较大的代码库中，建议降低此值，因为大量的控制流转换可能会增加代码的大小并降低代码的速度。
    deadCodeInjection: true, // 随机的死代码块(增加了混淆代码的大小)
    deadCodeInjectionThreshold: 0.1, // 死代码块的影响概率
    debugProtection: true, // 此选项几乎不可能使用开发者工具的控制台选项卡
    debugProtectionInterval: true, // 如果选中，则会在“控制台”选项卡上使用间隔强制调试模式，从而更难使用“开发人员工具”的其他功能。
    disableConsoleOutput: true, // 通过用空函数替换它们来禁用console.log，console.info，console.error和console.warn。这使得调试器的使用更加困难。
    identifierNamesGenerator: 'hexadecimal', // 标识符的混淆方式 hexadecimal(十六进制) mangled(短标识符)
    log: false,
    renameGlobals: false, // 是否启用全局变量和函数名称的混淆
    rotateStringArray: true, // 通过固定和随机（在代码混淆时生成）的位置移动数组。这使得将删除的字符串的顺序与其原始位置相匹配变得更加困难。如果原始源代码不小，建议使用此选项，因为辅助函数可以引起注意。
    selfDefending: true, // 混淆后的代码,不能使用代码美化,同时需要配置 compact:true;
    stringArray: true, // 删除字符串文字并将它们放在一个特殊的数组中
    stringArrayEncoding: ['rc4'],
    stringArrayThreshold: 1, // 您可以使用此设置调整字符串文字将插入stringArray的概率（从0到1）。此设置对于大型代码大小特别有用，因为它反复调用字符串数组并可能减慢代码速度
    transformObjectKeys: true,
    unicodeEscapeSequence: false// 允许启用/禁用字符串转换为unicode转义序列。Unicode转义序列大大增加了代码大小，并且可以轻松地将字符串恢复为原始视图。建议仅对小型源代码启用此选项。
  }
}

/**
 * @description 混淆处理，并将混淆的数据写到文件中
 * @param {Object}Element
 * @param {function}Callback
 */
const obfuscationProcess = (Element) => {
  let chooseConfig
  if (path.basename(Element.dirName).includes('app.')) {
    chooseConfig = obfuscationEncryptionConfig.max// 选择 最大压缩
  } else {
    chooseConfig = obfuscationEncryptionConfig.lower// 选择最低压缩
  }
  let encryption = JavaScriptObfuscator.obfuscate(Element.data.toString(), chooseConfig).getObfuscatedCode()// 进行混淆
  // 异步写出文件
  fs.writeFileSync(path.join(processPath, Element.dirName, path.sep), encryption)
}

let pb = null

function readFiles(FilePath) {
  return new Promise((res, rej) => {
    const actions = []
    const actions1 = []
    const fileList = []
    fs.readdir(FilePath, { withFileTypes: true }, (err, files) => {
      if (err) rej(err)
      if (files.length <= 0) return
      for (let index = 0; index < files.length; index++) {
        const file1 = files[index]
        const action = () => { // 将每一次循环方法定义为一个方法变量
          return new Promise(resolve => { // 每个方法返回一个Promise对象，第一个参数为resolve方法
            ((file) => {
              if (file.isFile()) {
                if (['.js'].includes(file.name.substr(-3))) {
                  // 是一个js文件
                  const child_filepath = `${FilePath}${file.name}`
                  fs.readFile(child_filepath, (err1, data) => {
                    if (err1) throw err1
                    const dir = '/' + FilePath.replace(processPath, '').replace('\\', '/')
                    const pushData = {
                      dir,
                      dirName: `${dir}${file.name}`,
                      data
                    }
                    fileList.push(pushData)
                    resolve()
                  })
                } else {
                  resolve()
                }
              } else {
                // 是一个目录
                const child_filepath = `${FilePath}${file.name}/`
                readFiles(child_filepath)
                resolve()
              }
            })(file1)
          })
        }
        actions.push(action())
      }
      // 调用Promise的all方法，传入方法数组，结束后执行then方法参数中的方法
      Promise.all(actions).then(() => {
        for (let index = 0; index < fileList.length; index++) {
          const action = () => {
            return new Promise(resolve => {
              const element = fileList[index]
              obfuscationProcess(element)
              pb.render({ completed: index, total: fileList.length, fileName: element.dirName })
              resolve()
            })
          }
          actions1.push(action())
        }
        Promise.all(actions1).then(() => {
          pb.render(
            {
              completed: fileList.length,
              total: fileList.length,
              fileDir: path.normalize(FilePath).replace(processPath, path.sep).replace(new RegExp(/\\/, 'g'), '/')
            }
          )
          res('开始进入这个目录下混淆')
          if (fileList.length > 0) { // 有操作文件，那么执行换行
            console.log('')// 不要删
          }
        })
      })
    })
  })
}

fs.readdir(processPath, { withFileTypes: true }, () => {
  console.log('正在混淆相关js代码包...')
  pb = new ProgressBar('正在混淆...', 0)
  readFiles(processPath)
});
// 监听销毁事件

['sigterm', 'exit'].forEach(Event => {
  process.on(Event, function() {
    console.log(`耗时时间...: ${((new Date() - processStartTime) / 1000).toFixed(2)}s`,)
    process.exit(0)
  })
})
