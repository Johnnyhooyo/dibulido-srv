package response

import "fmt"

// DibuError 自定义错误类型
type DibuError struct {
    Code CodeEnum
    Msg  MsgEnum
}

// Error toString方法 实现继承error
func (e DibuError) Error() string {
    return fmt.Sprintf("{\"code\":%s,\"msg\":%s}", string(rune(e.Code)), string(e.Msg))
}

var (
    SystemError       = DibuError{Code: SystemErrorCode, Msg: SystemErrorMsg}
    DBError           = DibuError{Code: DBErrorCode, Msg: DBErrorMsg}
    ParamError        = DibuError{Code: ParamErrorCode, Msg: ParamErrorMsg}
    PasswordIncorrect = DibuError{Code: PasswordIncorrectCode, Msg: PasswordIncorrectMSg}
    TokenExpired      = DibuError{Code: TokenExpiredCode, Msg: TokenExpiredMsg}
    TokenInvalid      = DibuError{Code: TokenInvalidCode, Msg: TokenInvalidMsg}
    UserNotExists     = DibuError{Code: UserNotExistsCode, Msg: UserNotExistsMsg}
    CasbinError       = DibuError{Code: CasbinOperatorErrorCode, Msg: CasbinOperatorErrorMsg}
    PermissionForbid  = DibuError{Code: PermissionForbidCode, Msg: PermissionForbidMsg}
)

type CodeEnum int
type MsgEnum string

const (
    SuccessCode     CodeEnum = 200
    SystemErrorCode CodeEnum = iota + 100000
    DBErrorCode
    ParamErrorCode
    PasswordIncorrectCode
    TokenExpiredCode
    TokenInvalidCode
    UserNotExistsCode
    CasbinOperatorErrorCode
    PermissionForbidCode
)

const (
    SuccessMsg             MsgEnum = "请求成功"
    SystemErrorMsg         MsgEnum = "系统错误"
    DBErrorMsg             MsgEnum = "数据库操作错误"
    ParamErrorMsg          MsgEnum = "请求参数错误"
    PasswordIncorrectMSg   MsgEnum = "密码错误"
    TokenExpiredMsg        MsgEnum = "登陆过期，请重新登录"
    TokenInvalidMsg        MsgEnum = "登陆异常，请稍后重试"
    UserNotExistsMsg       MsgEnum = "用户不存在，请先注册"
    CasbinOperatorErrorMsg MsgEnum = "权限操作失败，请稍后重试"
    PermissionForbidMsg    MsgEnum = "权限不足，请先申请权限"
)
