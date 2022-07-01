package user

import (
    "dibulido-srv/global"
    "dibulido-srv/global/log"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

// User 用户对象
type User struct {
    *gorm.Model
    UserID   string `json:"userID" gorm:"comment:用户ID"`                                          // 用户ID
    Username string `json:"userName" gorm:"comment:用户登录名" binding:"required" valid:"required"`   // 用户登录名
    Password string `json:"password"  gorm:"comment:用户登录密码" binding:"required" valid:"required"` // 用户登录密码
    NickName string `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                           // 用户昵称
    Phone    string `json:"phone"  gorm:"comment:用户手机号"`                                         // 用户手机号
    Email    string `json:"email"  gorm:"comment:用户邮箱"`                                          // 用户邮箱
    // AuthorityId string         `json:"authorityId" gorm:"default:888;comment:用户角色ID"` // 用户角色ID
    // Authority   SysAuthority   `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
    // Authorities []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
}

// Create 写入一个user到数据库
func (u *User) Create(c *gin.Context) error {
    result := global.DBMap[global.ConfigDB].Select("UserName").Create(u)
    log.WithContext(c).Info("保存用户信息结果:", result)
    return result.Error
}
