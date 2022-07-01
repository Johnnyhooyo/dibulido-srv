package model

// Response API响应结果
type Response struct {
    Code int         `json:"code"` // 结果码
    Msg  string      `json:"msg"`  // 结果描述
    Data interface{} `json:"data"` // 结果数据
}

const (
    RespSuccess = Response{Code: 200, Msg: "success"}
)
