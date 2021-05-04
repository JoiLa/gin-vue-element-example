package jwt

import (
    "net/http"
    "time"

    "Api/pkg/e"
    "Api/pkg/util"
    "github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
    return func(c *gin.Context) {
        var code int
        var data interface{}
        code = e.SUCCESS
        token := c.GetHeader("X-Token")
        if token == "" {
            code = e.InvalidParams
        } else {
            claims, err := util.ParseToken(token)
            if err != nil {
                code = e.ErrorAuthCheckTokenFail
            } else if time.Now().Unix() > claims.ExpiresAt {
                code = e.ErrorAuthCheckTokenTimeout
            }
        }
        if code != e.SUCCESS {
            c.JSON(http.StatusUnauthorized, gin.H{
                e.ParamsCode: code,
                e.ParamsMsg:  e.GetMsg(code),
                e.ParamsData: data,
            })
            c.Abort()
            return
        }

        c.Next()
    }
}
