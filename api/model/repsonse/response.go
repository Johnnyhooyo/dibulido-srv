package response

import (
    "dibulido-srv/global/log"
    "dibulido-srv/global/response"
    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "net/http"
)

// Response API响应结果
type Response struct {
    Code int         `json:"code"` // 结果码
    Msg  string      `json:"msg"`  // 结果描述
    Data interface{} `json:"data"` // 结果数据
}

// result 返回结果
func result(code response.CodeEnum, msg response.MsgEnum, data interface{}, c *gin.Context) {
    log.WithContext(c).Infof("请求结果:%d, %s, %+v", code, msg, data)
    // 开始时间
    c.JSON(http.StatusOK, Response{
        int(code),
        string(msg),
        data,
    })
}

// Success 返回空的成功请求结果
func Success(c *gin.Context) {
    result(response.SuccessCode, response.SuccessMsg, map[string]interface{}{}, c)
}

// SuccessWithData 返回带结果的成功请求
func SuccessWithData(data interface{}, c *gin.Context) {
    result(response.SuccessCode, response.SuccessMsg, data, c)
}

// FailWithError 根据制定的错误类型返回结果
func FailWithError(err error, c *gin.Context) {
    if dibuError, ok := err.(response.DibuError); ok {
        result(dibuError.Code, dibuError.Msg, map[string]interface{}{}, c)
    } else {
        Fail(c)
    }
}

// FailWithData 根据制定的错误类型返回结果
func FailWithData(err error, data interface{}, c *gin.Context) {
    if dibuError, ok := err.(response.DibuError); ok {
        result(dibuError.Code, dibuError.Msg, data, c)
    } else {
        Fail(c)
    }
}

// Fail 返回系统错误
func Fail(c *gin.Context) {
    result(response.SystemErrorCode, response.SystemErrorMsg, map[string]interface{}{}, c)
}

// ParamCheckFail 参数检查错误返回
func ParamCheckFail(err error, c *gin.Context) {
    switch err.(type) {
    case *validator.InvalidValidationError:
        FailWithData(response.ParamError, err.Error(), c)
        break
    case validator.ValidationErrors:
        FailWithData(response.ParamError, err.Error(), c)
        break
    default:
        FailWithError(response.ParamError, c)
    }
}
