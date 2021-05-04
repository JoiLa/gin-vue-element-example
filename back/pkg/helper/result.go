/**
 * description:
 * author: lfs <1144828910@qq.com>
 * Date: 2021/3/11
 */
package helper

import (
    `crypto/md5`
    `fmt`
    `time`

    `Api/pkg/e`
    `github.com/gin-gonic/gin`
)

type result struct {
    Ctx *gin.Context
}

func NewResult(ctx *gin.Context) *result {
    return &result{Ctx: ctx}
}

// 成功
func (r *result) Success(Code int, Msg string, Data interface{}) {
    response := gin.H{
        e.ParamsCode:      Code,
        e.ParamsMsg:       Msg,
        e.ParamsRequestId: fmt.Sprintf("%x", md5.Sum([]byte(time.Now().String()))),
    }
    if Data != nil {
        response[e.ParamsData] = Data
    }
    r.Ctx.Set(e.ParamsResponseJson, response)
}

// 出错
func (r *result) Error(Code int, Msg string) {
    r.Ctx.Set(e.ParamsResponseJson, gin.H{
        e.ParamsCode: Code,
        e.ParamsMsg:  Msg,
    })
}
