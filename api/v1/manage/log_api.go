package manage

import (
    response "dibulido-srv/api/model/repsonse"
    "dibulido-srv/api/v1/util"
    "dibulido-srv/dto/req"
    "dibulido-srv/global/log"
    "github.com/gin-gonic/gin"
)

// LogApi 登陆API
type LogApi struct{}

// Login 登陆接口
// @Description 登陆接口
// @Tags manage
// @Param login body req.LoginDto true "用户登陆参数"
// @Success 200 {object} response.Response{} "登陆成功"
// @Router /log/login [post]
func (api *LogApi) Login(c *gin.Context) {
    login := &req.LoginDto{}
    if success := util.BindParam(login, c); !success {
        return
    }
    log.WithContext(c).Infof("[Login] params:%+v", login)

    if err := logService.Login(login, c); err != nil {
        response.FailWithError(err, c)
    } else {
        response.Success(c)
    }
}

// Logout 登出接口
// @Description 登出接口
// @Tags manage
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{} "登出成功"
// @Router /log/logout [post]
func (api *LogApi) Logout(c *gin.Context) {
    if err := logService.Logout(c); err != nil {
        response.FailWithError(err, c)
    } else {
        response.Success(c)
    }
}
