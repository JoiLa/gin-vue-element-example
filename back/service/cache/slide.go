/**
 * description:
 * author: lfs <1144828910@qq.com>
 * Date: 2021/3/7
 */
package cache

import (
    `encoding/json`
    `errors`
    `time`
)

var prefix = "slide::"

var slideAnswer map[string][]byte // 保存滑动位置的答案

func init() {
    slideAnswer = make(map[string][]byte)
}

/**
 * @Description: 保存滑动答案
 * @param Seed
 * @param Position
 */
func SaveSlideAnswer(Seed string, Position interface{}) {
    // cache.RedisDb.Set(prefix+Seed, bytes, time.Second*120) // 使用redis缓存
    go func() {
        bytes, _ := json.Marshal(Position)
        slideAnswer[prefix+Seed] = bytes
        for i := 0; i < 120; i++ {
            time.Sleep(time.Second * 1)
        }
        delete(slideAnswer, prefix+Seed)
    }()
}

/**
 * @Description: 获取滑动答案
 * @param Seed
 * @param Position
 */
/*func GetSlideValue(Seed string) (*redis.StringCmd, error) {
    read := cache.RedisDb.Get(prefix + Seed)
      err := read.Err()
      if err != nil {
          return nil, err
      }
    return read, nil
}*/
/**
 * @Description: 获取滑动答案
 * @param Seed
 * @param Position
 */
func GetSlideValue(Seed string) ([]byte, error) {
    read := slideAnswer[prefix+Seed]
    if read == nil {
        return nil, errors.New("cache data not exist")
    }
    return read, nil
}
