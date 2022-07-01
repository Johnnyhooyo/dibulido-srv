package manage

import (
    v1 "dibulido-srv/api/v1"
    "dibulido-srv/middleware"
    "github.com/gin-gonic/gin"
)

// LogRouter 登陆相关操作接口
type LogRouter struct{}

// InitILogRouter 登陆相关操作接口
func (s *LogRouter) InitILogRouter(Router *gin.RouterGroup) {
    logAPI := v1.ApiGroupApp.ManageGroup.LogApi

    traceLogRouter := Router.Group("log").Use(middleware.TraceLoggerMiddleware())
    {
        traceLogRouter.POST("login", logAPI.Login)
    }

    // 需要登陆的接口
    loginTraceLogRouter := Router.Group("log").
        Use(middleware.TraceLoggerMiddleware()).
        Use(middleware.JwtMiddleware())
    {
        loginTraceLogRouter.POST("logout", logAPI.Logout)
    }
}
