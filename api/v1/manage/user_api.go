package manage

import (
    "dibulido-srv/api/model/repsonse"
    "dibulido-srv/api/v1/util"
    "dibulido-srv/dto/req"
    "dibulido-srv/global/log"
    "github.com/gin-gonic/gin"
)

// IUserApi 用户API
type IUserApi struct{}

// RegisterUser 注册新用户
// @Description 注册新用户接口
// @Tags manage
// @Param userCreate body req.UserCreateDto true "用户参数"
// @Success 200 {object} response.Response{} "注册用户成功"
// @Router /user/register [post]
func (api *IUserApi) RegisterUser(c *gin.Context) {
    u := &req.UserCreateDto{}
    if success := util.BindParam(u, c); !success {
        return
    }
    log.WithContext(c).Infof("[RegisterUser] params:%+v", u)

    if userID, err := userService.RegisterUser(u, c); err != nil {
        response.FailWithError(err, c)
    } else {
        response.SuccessWithData(userID, c)
    }
}

// ModifyUser 用户信息变更
// @Description 用户信息变更
// @Tags manage
// @Security ApiKeyAuth
// @Param userUpdate body req.UserUpdateDto true "用户参数"
// @Success 200 {object} response.Response{} "修改成功"
// @Router /user/modify [post]
func (api *IUserApi) ModifyUser(c *gin.Context) {
    u := &req.UserUpdateDto{}
    if success := util.BindParam(u, c); !success {
        return
    }
    log.WithContext(c).Infof("[RegisterUser] params:%+v", u)

    if userID, err := userService.UpdateUser(u, c); err != nil {
        response.FailWithError(err, c)
    } else {
        response.SuccessWithData(userID, c)
    }
}

// CloseUser 注销账号
// @Description 注销账号
// @Tags manage
// @Security ApiKeyAuth
// @Param userClose body req.UserCloseDto true "用户ID"
// @Success 200 {object} response.Response{} "注销成功"
// @Router /user/close [post]
func (api *IUserApi) CloseUser(c *gin.Context) {
    u := &req.UserCloseDto{}
    if success := util.BindParam(u, c); !success {
        return
    }
    log.WithContext(c).Infof("[CloseUser] params:%+v", u)

    if err := userService.DelUser(u, c); err != nil {
        response.FailWithError(err, c)
    } else {
        response.Success(c)
    }
}
