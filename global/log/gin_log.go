package log

import (
    "dibulido-srv/global"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

// 在gin.Context中日志对象的key
const loggerKey = "ginLogger"

// NewContext 给gin content设置新的值
func NewContext(ctx *gin.Context, fields ...zapcore.Field) {
    log := WithContext(ctx)
    for _, field := range fields {
        log = log.With(field)
    }
    ctx.Set(loggerKey, log)
}

// WithContext 获取对应gin context的SugaredLogger 如果不存在获取的是全局SugaredLogger
func WithContext(ctx *gin.Context) *zap.SugaredLogger {
    if ctx == nil {
        return global.Log.SugaredLogger
    }
    log, _ := ctx.Get(loggerKey)
    if ctxLogger, ok := log.(*zap.SugaredLogger); ok {
        return ctxLogger
    }
    return global.Log.SugaredLogger
}
