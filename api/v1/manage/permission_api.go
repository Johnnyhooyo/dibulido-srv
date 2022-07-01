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
// @Security ApiKeyAuth
// @Param userRole body common.UserRoleInfo true "用户权限"
// @Success 200 {object} response.Response{} "操作成功"
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

// RemoveRoleFromUser 取消用户已授权角色
// @Description 取消用户已授权角色
// @Tags manage
// @Security ApiKeyAuth
// @Param userRole body common.UserRoleInfo true "取消用户授权"
// @Success 200 {object} response.Response{} "操作成功"
// @Router /pms/remove [post]
func (p *PermissionAPI) RemoveRoleFromUser(c *gin.Context) {
    param := &common.UserRoleInfo{}
    if success := util.BindParam(param, c); !success {
        return
    }
    log.WithContext(c).Infof("[RemoveRoleFromUser] params:%+v", param)

    for _, ID := range param.RoleIDs {
        if err := casbinService.DeleteRoleForUser(param.UserID, ID); err != nil {
            response.FailWithError(err, c)
            return
        }
    }
    response.Success(c)
}

// GetUserRoles 获取用户已授权角色
// @Description 获取用户已授权角色
// @Tags manage
// @Security ApiKeyAuth
// @Param userID query string true "用户ID"
// @Success 200 {object} response.Response{} "操作成功"
// @Router /pms/query [get]
func (p *PermissionAPI) GetUserRoles(c *gin.Context) {
    userID := c.Query("userID")
    log.WithContext(c).Infof("[GetUserRoles] params:%s", userID)

    if roleInfo, err := casbinService.GetRolesForUser(userID); err != nil {
        response.FailWithError(err, c)
        return
    } else {
        response.SuccessWithData(roleInfo, c)
    }
}
