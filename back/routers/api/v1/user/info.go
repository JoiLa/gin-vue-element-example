/**
 * description:
 * author: lfs <1144828910@qq.com>
 * Date: 2021/5/3
 */
package user

import (
    `Api/pkg/e`
    `Api/pkg/helper`
    `github.com/gin-gonic/gin`
)

type userInfo struct {
    Roles        []string `json:"roles"`
    Introduction string   `json:"introduction"`
    Avatar       string   `json:"avatar"`
    Name         string   `json:"name"`
}

func Info(c *gin.Context) {
    var temp userInfo
    temp.Roles = []string{"admin"}
    temp.Introduction = "I am a super administrator"
    temp.Avatar = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
    temp.Name = "Super Admin"
    // 滑动到阈值范围
    helper.NewResult(c).Success(e.SUCCESS, e.GetMsg(e.SUCCESS), temp)
}
