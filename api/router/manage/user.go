package user

import (
    v1 "dibulido-srv/api/v1"
    "dibulido-srv/middleware"
    "github.com/gin-gonic/gin"
)

// IUserRouter 用户相关操作接口
type IUserRouter struct{}

// InitIUserRouter 用户相关操作接口
func (s *IUserRouter) InitIUserRouter(Router *gin.RouterGroup) {
    manageRouter := Router.Group("manage")

    traceUserRouter := manageRouter.Group("user").Use(middleware.TraceLoggerMiddleware())
    {
        userAPI := v1.ApiGroupApp.UserGroup.IUserApi
        traceUserRouter.POST("register", userAPI.RegisterUser)
        traceUserRouter.POST("modify", userAPI.ModifyUser)
        traceUserRouter.POST("close", userAPI.CloseUser)
    }

}
