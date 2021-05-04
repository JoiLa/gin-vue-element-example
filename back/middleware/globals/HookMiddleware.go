package globals

import (
    "github.com/gin-gonic/gin"
)

// bodyLogWriter是为了记录返回数据到log中进行了双写
/**
 * @Description: 全局监听中间件
 * @return gin.HandlerFunc
 */
func HookMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 请求之前
        // fmt.Println(c.ClientIP(), "请求之前")
        c.Next()

        // 请求之后
        // fmt.Println(c.ClientIP(), "请求之后")
    }
}
