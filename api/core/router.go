package router

import (
    v1 "dibulido-srv/api/v1"
    "github.com/gin-gonic/gin"
)

func InitGinRouter() *gin.Engine {
    r := gin.Default()
    backendRouter := v1.RouterGroupApp.Backend
    baseRouter := r.Group("")
    {
        backendRouter.InitBackendRouter(baseRouter)
    }

    return r
}
