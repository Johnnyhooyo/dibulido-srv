package manage

import (
    "dibulido-srv/global"
    "dibulido-srv/global/log"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

// User 用户对象
type User struct {
    *gorm.Model
    UserID   string `json:"userID" gorm:"comment:用户ID"`      // 用户ID
    UserName string `json:"userName" gorm:"comment:用户登录名"`   // 用户登录名
    Password string `json:"password"  gorm:"comment:用户登录密码"` // 用户登录密码
    NickName string `json:"nickName" gorm:"comment:用户昵称"`    // 用户昵称
    Phone    string `json:"phone"  gorm:"comment:用户手机号"`     // 用户手机号
    Email    string `json:"email"  gorm:"comment:用户邮箱"`      // 用户邮箱
}

// TableName 用户表表名称
func (u *User) TableName() string {
    return "d_user_base_info"
}

// Create 写入一个user到数据库
func (u *User) Create(c *gin.Context) error {
    result := global.DBMap[global.ConfigDB].Create(u)
    log.WithContext(c).Debugf("保存用户信息结果:%d", result.RowsAffected)
    return result.Error
}

// Update 更新用户信息
func (u *User) Update(c *gin.Context) error {
    result := global.DBMap[global.ConfigDB].Where("user_id=?", u.UserID).Updates(u)
    log.WithContext(c).Debugf("更新用户信息结果:%+v", result)
    return result.Error
}

// Del 删除用户
func (u *User) Del(c *gin.Context) error {
    result := global.DBMap[global.ConfigDB].Where("user_id=?", u.UserID).Delete(u)
    log.WithContext(c).Debugf("更新用户信息结果:%+v", result)
    return result.Error
}

// QueryByName 根据登陆名成查询用户信息
func (u *User) QueryByName(c *gin.Context) (*User, error) {
    user := &User{}
    result := global.GetDB(global.ConfigDB, c).Where("user_name=?", u.UserName).First(user)
    log.WithContext(c).Debugf("获取用户信息结果:%+v", result)
    return user, result.Error
}
