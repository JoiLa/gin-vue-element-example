/**
 * description:
 * author: lfs <1144828910@qq.com>
 * Date: 2021/3/7
 */
package cache

import (
    `log`

    `Api/pkg/setting`
    "github.com/go-redis/redis"
)

var (
    RedisDb *redis.Client // 保存redis连接
)

// 创建到redis的连接
func init1() {
    sec, err := setting.Cfg.GetSection("redis")
    if err != nil {
        log.Fatal(2, "Fail to get section 'redis': %v", err)
    }
    RedisDb = redis.NewClient(&redis.Options{
        Addr:     sec.Key("HOST").String(),
        Password: sec.Key("PASSWORD").String(), // no password set
        DB:       0,                            // use default DB
    })
    _, err = RedisDb.Ping().Result()
    if err != nil {
        log.Fatal(2, "Fail to connection 'redis': %v", err)
    }
}
