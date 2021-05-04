/**
 * description:
 * author: lfs <1144828910@qq.com>
 * Date: 2021/3/5
 */
package helper

import (
    `math/rand`
    `os`
    `path/filepath`
    `time`

    `github.com/gobuffalo/packr`
)

type PackerStruct struct {
    StorageSide packr.Box
    StorageWeb  packr.Box
}

var AppBox PackerStruct

/**
 * @Description: 初始化
 */
func init() {
    AppBox.StorageSide = packr.NewBox("../../public/storage/side")
    AppBox.StorageWeb = packr.NewBox("../../public/web")
}

/**
 * @Description: 获取应用路径
 * @param Path
 * @return string
 */
func AppPath(Path string) string {
    pwd, _ := os.Getwd() // 获取当前运行的目录
    return filepath.Clean(filepath.Join(pwd, Path))
}

/**
 * @Description: 获取主路径
 * @param Path
 * @return string
 */
func PublicPath(Path string) string {
    return filepath.Clean(filepath.Join(AppPath("public"), Path))
}

/**
 * @Description: 获取存储路径
 * @param Path
 * @return string
 */
func StoragePath(Path string) string {
    return filepath.Clean(filepath.Join(PublicPath("storage"), Path))
}

/**
 * @Description: 随机整数
 * @param min
 * @param max
 * @return int
 */
func RandInt(min, max int) int {
    if min >= max || min == 0 || max == 0 {
        return max
    }
    rand.Seed(time.Now().UnixNano()) // 将时间戳设置成种子数
    max++
    return rand.Intn(max-min) + min
}

/**
 * @Description: 随机生成float64
 * @param min
 * @param max
 * @return int
 */
func RandFloat64(min, max float64) float64 {
    if min >= max || min == 0 || max == 0 {
        return max
    }
    rand.Seed(time.Now().UnixNano()) // 将时间戳设置成种子数
    return min + rand.Float64()*(max-min)
}
