package backend

import (
    v1 "dibulido-srv/api/v1"
    "dibulido-srv/middleware"
    "github.com/gin-gonic/gin"
)

// HealthRouter 健康检查操作
type HealthRouter struct{}

// InitHealthRouter 初始化健康检查
func (s *HealthRouter) InitHealthRouter(Router *gin.RouterGroup) {
    backendRouter := Router.Group("health")

    traceBackendRouter := backendRouter.Use(middleware.TraceLoggerMiddleware())
    {
        healthApi := v1.ApiGroupApp.BackendGroup.HealthApi
        traceBackendRouter.GET("ping", healthApi.Ping)
    }

}
