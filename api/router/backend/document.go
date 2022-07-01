package backend

import (
    "dibulido-srv/docs"
    "dibulido-srv/middleware"
    "github.com/gin-gonic/gin"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)

// DocRouter 文档路由
type DocRouter struct{}

// InitDocRouter 初始化文档路由
func (s *DocRouter) InitDocRouter(Router *gin.RouterGroup) {
    backendRouter := Router.Group("doc")
    {
        // swagger路由 需要登陆后查看
        // swagger路径： /doc/swagger/index.html
        swaggerRouter := backendRouter.Use(middleware.TraceLoggerMiddleware())
        // 如果swagger报错500 很可能是这里路径不对
        docs.SwaggerInfo.BasePath = "/"
        swaggerRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    }
}
