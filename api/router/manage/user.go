package manage

import (
    v1 "dibulido-srv/api/v1"
    "dibulido-srv/middleware"
    "github.com/gin-gonic/gin"
)

// IUserRouter 用户相关操作接口
type IUserRouter struct{}

// InitIUserRouter 用户相关操作接口
func (s *IUserRouter) InitIUserRouter(Router *gin.RouterGroup) {
    userAPI := v1.ApiGroupApp.ManageGroup.IUserApi

    traceUserRouter := Router.Group("user").Use(middleware.TraceLoggerMiddleware())
    {

        traceUserRouter.POST("register", userAPI.RegisterUser)

    }

    // 需要登陆的接口
    loginTraceUserRouter := Router.Group("user").
        Use(middleware.TraceLoggerMiddleware()).
        Use(middleware.JwtMiddleware()).
        Use(middleware.CasbinMiddleware())
    {
        loginTraceUserRouter.POST("modify", userAPI.ModifyUser)
        loginTraceUserRouter.POST("close", userAPI.CloseUser)
    }

}
