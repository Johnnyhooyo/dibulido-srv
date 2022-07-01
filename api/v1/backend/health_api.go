package backend

import (
    "dibulido-srv/global/log"
    "github.com/gin-gonic/gin"
)

// HealthApi 健康检测API
type HealthApi struct{}

// Ping 健康检查
// @Tags backends
// @Summary 健康检查
// @Success 200 "pong"
// @Router /health/ping [get]
func (s *HealthApi) Ping(c *gin.Context) {
    log.WithContext(c).Info("ping")
    c.JSON(200, "pong")
}
