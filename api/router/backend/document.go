package beckend

import (
    "dibulido-srv/middleware"
    "github.com/gin-contrib/pprof"
    "github.com/gin-gonic/gin"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    "github.com/swaggo/swag/example/basic/docs"
)

// BackendRouter 后台操作路由
type BackendRouter struct{}

// InitBackendRouter 初始化后台操作路由
func (s *BackendRouter) InitBackendRouter(Router *gin.RouterGroup) {
    backendRouter := Router.Group("backend")
    {
        // swagger 路由 需要登陆后查看 todo middleware替换成登陆
        // swagger 路径： /dibulido/backend/swagger/index.html
        swaggerRouter := backendRouter.Use(middleware.TraceLoggerMiddleware())
        docs.SwaggerInfo.BasePath = "/dibulido/backend"
        swaggerRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    }
    {
        // 注册pprof性能分析组件 todo 这里看下是否增加一个权限中间件 不能随便访问
        // pprof 路径： /dibulido/backend/debug/pprof
        pprof.RouteRegister(backendRouter)
    }

}
