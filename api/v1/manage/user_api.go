package user

import (
    "dibulido-srv/api/model/repsonse"
    "dibulido-srv/global/log"
    "dibulido-srv/model/user"
    "github.com/gin-gonic/gin"
    "github.com/gofrs/uuid"
)

// IUserApi 用户API
type IUserApi struct{}

// RegisterUser 注册新用户
// @Description 注册新用户接口
// @Tags manage
// @Param manage body manage.User true "用户参数"
// @Success 200 {object} response.Response{} "注册用户成功"
// @Router /manage/register [post]
func (api *IUserApi) RegisterUser(c *gin.Context) {
    u := &user.User{}
    err := c.ShouldBindJSON(u)
    if err != nil {
        log.WithContext(c).Errorf("获取请求参数失败,%s", err.Error())
        response.FailWithError(response.ParamError, c)
        return
    }
    if uuidV4, err := uuid.NewV4(); err != nil {
        log.WithContext(c).Errorf("RegisterUser get uuidv4 error:%s", err.Error())
        response.Fail(c)
        return
    } else {
        u.UserID = uuidV4.String()
    }
    if err := u.Create(c); err != nil {
        log.WithContext(c).Errorf("创建用户失败，%s", err.Error())
        response.FailWithMsg(response.DBErrorCode, response.DBErrorMsg, c)
        return
    }
    response.Success(c)
}
