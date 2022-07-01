package req

// User 用户公用属性
type User struct {
    // max length: 64
    NickName string `json:"nickName" binding:"omitempty,lt=64"`
    // max length: 20
    Phone string `json:"phone" binding:"omitempty,lt=20"`
    // max length: 64
    Email string `json:"email" binding:"omitempty,lt=64,email"`
}

// UserUpdateDto 修改用户请求参数
// swagger:model userUpdate
type UserUpdateDto struct {
    UserID string `json:"userID" binding:"required"`
    // max length: 32
    UserName string `json:"userName" binding:"omitempty,lt=32"`
    // max length: 32
    Password string `json:"password" binding:"omitempty,lt=32"`
    User
}

// UserCreateDto 创建用户请求参数
// swagger:model userCreate
type UserCreateDto struct {
    // max length: 32
    UserName string `json:"userName" binding:"required,lt=32"`
    // max length: 32
    Password string `json:"password" binding:"required,lt=32"`
    User
}

// UserCloseDto 注销用户请求
// swagger:model userClose
type UserCloseDto struct {
    UserID string `json:"userID" binding:"required"`
}
