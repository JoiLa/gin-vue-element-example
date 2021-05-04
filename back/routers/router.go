package routers

import (
    "Api/middleware/globals"
    `Api/middleware/jwt`
    `Api/pkg/helper`
    "Api/pkg/setting"
    `Api/routers/api`
    `Api/routers/api/v1/slide`
    `Api/routers/api/v1/user`
    "github.com/gin-gonic/gin"
)

// 初始化路由
func InitRouter() *gin.Engine {

    r := gin.New()

    r.Use()
    {
        r.StaticFS("/", helper.AppBox.StorageWeb) // 挂载 静态资源
    }
    r.Use(globals.HookMiddleware(), globals.ParseCiphertextMiddleware(), globals.GenerateCiphertextMiddleware(), gin.Logger(), gin.Recovery())

    gin.SetMode(setting.RunMode)

    r.POST("/auth", api.GetAuth) // 登录板块

    side := r.Group("/side") // 滑动验证板块
    side.Use()
    {
        // 生成滑动图片
        side.POST(`/generate-image`, slide.GenerateImage)
        // 效验滑动是否正确
        side.POST(`/verify-slide`, slide.VerifySlide)
    }

    apiV1 := r.Group("/v1")
    apiV1.Use(jwt.JWT())
    {
        apiV1.POST("/user/info", user.Info)     // 用户信息
        apiV1.POST("/user/logout", user.Logout) // 用户退出
    }
    return r
}
