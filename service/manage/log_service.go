package manage

import (
    "dibulido-srv/dto/req"
    "dibulido-srv/global"
    "dibulido-srv/global/log"
    "dibulido-srv/global/response"
    "dibulido-srv/model/manage"
    "github.com/gin-gonic/gin"
    "strconv"
)

// LogService 登陆服务
type LogService struct{}

// Login 用户登陆
func (l *LogService) Login(loginDto *req.LoginDto, c *gin.Context) error {
    u := &manage.User{
        UserName: loginDto.UserName,
    }
    if user, err := u.QueryByName(c); err != nil {
        return response.UserNotExists
    } else {
        if loginDto.Password != user.Password {
            return response.PasswordIncorrect
        }
        userInfo := &global.UserInfo{
            UserID:   user.UserID,
            UserName: user.UserName,
        }
        token, err := global.Sign(userInfo)
        if err != nil {
            log.WithContext(c).Errorf("用户:%s 登陆失败(签名失败):%s", loginDto.UserName, err.Error())
            return err
        }
        newClaims, err := global.VerifySign(token)
        if err != nil {
            log.WithContext(c).Errorf("用户:%s 登陆失败:%s", loginDto.UserName, err.Error())
            return err
        }
        c.Header(global.TokenHeaderKey, token)
        c.Header(global.TokenExpiredKey, strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
    }
    return nil
}

// Logout 用户登出
func (l *LogService) Logout(c *gin.Context) error {
    userInfo, _ := c.Get(global.GinUserInfoKey)
    log.WithContext(c).Infof("用户登出:%+v", userInfo)
    // todo 后续可以加上redis 来做登录态剔除
    return nil
}
