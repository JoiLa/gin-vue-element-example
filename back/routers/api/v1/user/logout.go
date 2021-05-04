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

func Logout(c *gin.Context) {
    // 滑动到阈值范围
    helper.NewResult(c).Success(e.SUCCESS, e.GetMsg(e.SUCCESS), nil)
}
