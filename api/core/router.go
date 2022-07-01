package core

import (
    "dibulido-srv/api/router"
    "github.com/gin-gonic/gin"
)

// InitGinRouter 初始化Gin
func InitGinRouter() *gin.Engine {
    r := gin.Default()
    baseRouter := r.Group("")

    backendRouter := router.BaseRouterGroupApp.Backend
    {
        backendRouter.InitPPRofRouter(baseRouter)  // 性能剖析组件
        backendRouter.InitHealthRouter(baseRouter) // 健康检查接口
        backendRouter.InitDocRouter(baseRouter)    // swagger文档接口
    }

    manageRouter := router.BaseRouterGroupApp.Manage
    {
        manageRouter.InitIUserRouter(baseRouter)       // 注册用户接口
        manageRouter.InitILogRouter(baseRouter)        // 登陆相关接口
        manageRouter.InitIPermissionRouter(baseRouter) // 权限相关接口
    }

    return r
}
