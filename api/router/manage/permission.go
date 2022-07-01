package manage

import (
    v1 "dibulido-srv/api/v1"
    "dibulido-srv/middleware"
    "github.com/gin-gonic/gin"
)

// PermissionRouter 权限相关操作接口
type PermissionRouter struct{}

// InitIPermissionRouter 权限相关操作接口
// 需要管理员权限
func (s *PermissionRouter) InitIPermissionRouter(Router *gin.RouterGroup) {
    permissionRouter := Router.Group("pms").
        Use(middleware.TraceLoggerMiddleware()).
        Use(middleware.JwtMiddleware()).
        Use(middleware.CasbinMiddleware())

    permissionAPI := v1.ApiGroupApp.ManageGroup.PermissionAPI
    {
        permissionRouter.POST("grant", permissionAPI.GrantUserRole)
        permissionRouter.POST("remove", permissionAPI.RemoveRoleFromUser)
        permissionRouter.GET("query", permissionAPI.GetUserRoles)
    }
}
