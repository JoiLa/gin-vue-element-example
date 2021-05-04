package e

var MsgFlags = map[int]string{
    SUCCESS:                              "成功",
    ERROR:                                "错误",
    InvalidParams:                        "请求参数错误",
    VerificationFailedPleaseTryAgain:     "验证失败，请重试",
    VerificationHasFailedPleaseReAcquire: "验证已经时效，请重新获取",
    VerificationFrequentFailedPleaseReAcquire: "验证太频繁，请重新获取",
    ErrorAuthCheckTokenFail:                   "Token鉴权失败",
    ErrorAuthCheckTokenTimeout:                "Token已超时",
    ErrorAuthToken:                            "Token生成失败",
    ErrorAuth:                                 `账号或者密码错误`,
}

func GetMsg(code int) string {
    msg, ok := MsgFlags[code]
    if ok {
        return msg
    }

    return MsgFlags[ERROR]
}
