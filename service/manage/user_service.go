package manage

import (
    "dibulido-srv/dto/req"
    "dibulido-srv/global"
    "dibulido-srv/global/log"
    "dibulido-srv/global/response"
    "dibulido-srv/model/manage"
    "github.com/gin-gonic/gin"
    "github.com/gofrs/uuid"
)

// UserService 用户服务
type UserService struct{}

// RegisterUser 注册用户方法
func (userService *UserService) RegisterUser(user *req.UserCreateDto, c *gin.Context) (string, error) {
    u := &manage.User{}
    if err := global.DepCopy(user, u); err != nil {
        log.WithContext(c).Errorf("[RegisterUser] data copy error:%s", err.Error())
        return "", err
    }
    // 生成唯一ID
    if uuidV4, err := uuid.NewV4(); err != nil {
        log.WithContext(c).Errorf("RegisterUser get uuidv4 error:%s", err.Error())
        return "", err
    } else {
        u.UserID = uuidV4.String()
    }
    // 保存
    if err := u.Create(c); err != nil {
        log.WithContext(c).Errorf("创建用户失败，%s", err.Error())
        return "", response.DBError
    }
    return u.UserID, nil
}

// UpdateUser 修改用户信息方法
func (userService *UserService) UpdateUser(user *req.UserUpdateDto, c *gin.Context) (string, error) {
    u := &manage.User{}
    if err := global.DepCopy(user, u); err != nil {
        log.WithContext(c).Errorf("[UpdateUser] data copy error:%s", err.Error())
        return "", err
    }
    if err := u.Update(c); err != nil {
        log.WithContext(c).Errorf("修改用户失败，%s", err.Error())
        return "", response.DBError
    }
    return u.UserID, nil
}

// DelUser 删除用户
func (userService UserService) DelUser(user *req.UserCloseDto, c *gin.Context) error {
    u := &manage.User{
        UserID: user.UserID,
    }
    if err := u.Del(c); err != nil {
        log.WithContext(c).Errorf("删除用户失败，%s", err.Error())
        return response.DBError
    }
    return nil
}
