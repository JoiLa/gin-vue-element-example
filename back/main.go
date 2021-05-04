package main

import (
    "fmt"
    "net/http"
    `os/exec`
    `syscall`
    `time`

    "Api/pkg/setting"
    "Api/routers"
)

/**
 * @Description: 主要的入口
 */
func main() {
    router := routers.InitRouter()
    s := &http.Server{
        Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
        Handler:        router,
        ReadTimeout:    setting.ReadTimeout,
        WriteTimeout:   setting.WriteTimeout,
        MaxHeaderBytes: 1 << 20,
    }
    go func() {
        uri := fmt.Sprintf("http://127.0.0.1:%d", setting.HTTPPort)
        fmt.Println("即将在3秒后，打开网页：", uri)
        time.Sleep(time.Second * 3)
        err := Open(uri)
        if err != nil {
            fmt.Println("打开浏览器失败")
        }
    }()
    _ = s.ListenAndServe()
}

/**
 * @Description: 打开网页
 * @param uri
 * @return error
 */
func Open(uri string) error {
    cmd := exec.Command("cmd", "/c", "start", uri)
    cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
    return cmd.Start()
}
