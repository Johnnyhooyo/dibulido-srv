package middleware

import (
    "dibulido-srv/global"
    "dibulido-srv/global/log"
    "github.com/gin-gonic/gin"
    "github.com/gofrs/uuid"
    "go.uber.org/zap"
)

// TraceLoggerMiddleware zap日志中增加traceID
func TraceLoggerMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        traceID := c.GetHeader(global.TraceKey)
        if traceID == "" {
            if uuidV4, err := uuid.NewV4(); err != nil {
                global.Log.Errorf("TraceLoggerMiddleware get uuidv4 error:%s", err.Error())
            } else {
                traceID = uuidV4.String()
            }
        }
        // 如果有其他需要在请求级别的日志参数 可以在此设置
        c.Set(global.TraceKey, traceID)
        log.NewContext(c, zap.String(global.TraceKey, traceID))

        c.Next()
    }
}
