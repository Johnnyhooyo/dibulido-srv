package manage

import (
    response "dibulido-srv/api/model/repsonse"
    "dibulido-srv/api/v1/util"
    "dibulido-srv/dto/common"
    "dibulido-srv/global/log"
    "github.com/gin-gonic/gin"
)

// PermissionAPI 权限相关API
type PermissionAPI struct{}

// GrantUserRole 给用户授权角色
// @Description 给用户授权角色
// @Tags manage
// @Param userRole body common.UserRoleInfo true "用户权限"
// @Success 200 {object} response.Response{} "登陆成功"
// @Router /pms/grant [post]
func (p *PermissionAPI) GrantUserRole(c *gin.Context) {
    param := &common.UserRoleInfo{}
    if success := util.BindParam(param, c); !success {
        return
    }
    log.WithContext(c).Infof("[GrantUserRole] params:%+v", param)

    if err := casbinService.AddRolesForUser(param.UserID, param.RoleIDs...); err != nil {
        response.FailWithError(err, c)
    } else {
        response.Success(c)
    }
}
