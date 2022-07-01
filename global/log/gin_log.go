package global

import (
    "github.com/gin-gonic/gin"
    "go.uber.org/zap/zapcore"
    "strconv"
)

// 在gin.Context中日志对象的key
const loggerKey = iota

// NewContent 给gin content设置新的值
func NewContent(ctx *gin.Context, fields ...zapcore.Field) {
    ctx.Set(strconv.Itoa(loggerKey), WithContent(ctx).With(fields))
}

// WithContent 获取对应gin context的Logger 如果不存在获取的是全局Logger
func WithContent(ctx *gin.Context) *Logger {
    if ctx == nil {
        return Log
    }
    log, _ := ctx.Get(strconv.Itoa(loggerKey))
    if ctxLogger, ok := log.(*Logger); ok {
        return ctxLogger
    }
    return Log
}
