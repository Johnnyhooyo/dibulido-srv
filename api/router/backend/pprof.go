package backend

import (
    "github.com/gin-contrib/pprof"
    "github.com/gin-gonic/gin"
)

// PPRofRouter 性能检测路由
type PPRofRouter struct{}

// InitPPRofRouter 初始化性能检测路由
func (s *PPRofRouter) InitPPRofRouter(Router *gin.RouterGroup) {
    backendRouter := Router.Group("backend")

    {
        // 注册pprof性能分析组件
        // pprof 路径： /backend/debug/pprof
        pprof.RouteRegister(backendRouter)
    }

}
