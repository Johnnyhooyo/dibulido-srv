package global

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
    DBError    = DibuError{Code: DBErrorCode, Msg: DBErrorMsg}
    ParamError = DibuError{Code: ParamErrorCode, Msg: ParamErrorMsg}
)

type CodeEnum int
type MsgEnum string

const (
    SuccessCode     CodeEnum = 200
    SystemErrorCode CodeEnum = iota + 100001
    DBErrorCode
    ParamErrorCode
)

const (
    SuccessMsg     MsgEnum = "请求成功"
    SystemErrorMsg MsgEnum = "系统错误"
    DBErrorMsg     MsgEnum = "数据库操作错误"
    ParamErrorMsg  MsgEnum = "请求参数错误"
)
