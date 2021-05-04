package api

import (
    `Api/pkg/e`
    `Api/pkg/helper`
    `Api/pkg/logging`
    `Api/pkg/util`
    `github.com/astaxie/beego/validation`
    "github.com/gin-gonic/gin"
)

type auth struct {
    Username string `valid:"Required; MaxSize(50)"`
    Password string `valid:"Required; MaxSize(50)"`
}

/**
 * @Description: 认证用户
 * @param c
 */
func GetAuth(c *gin.Context) {
    username := c.Request.PostForm.Get(e.ParamsUsername)
    password := c.Request.PostForm.Get(e.ParamsPassword)
    valid := validation.Validation{}
    a := auth{Username: username, Password: password}
    ok, _ := valid.Valid(&a)

    data := make(map[string]interface{})
    code := e.InvalidParams
    if ok {
        if username == "admin" && password == "123456" {
            token, err := util.GenerateToken(username)
            if err != nil {
                code = e.ErrorAuthToken
            } else {
                data[e.ParamsToken] = token
                code = e.SUCCESS
            }
        } else {
            code = e.ErrorAuth
        }
    } else {
        for _, err := range valid.Errors {
            logging.Info(err.Key, err.Message)
        }
    }
    helper.NewResult(c).Success(code, e.GetMsg(code), data)
}
