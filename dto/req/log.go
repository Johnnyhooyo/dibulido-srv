package req

// LoginDto 用户登陆请求参数
// swagger:model login
type LoginDto struct {
    // max length: 32
    UserName string `json:"userName" binding:"required,lt=32"`
    // max length: 32
    Password string `json:"password" binding:"required,lt=32"`
}
