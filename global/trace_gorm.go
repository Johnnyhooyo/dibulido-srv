package global

import (
    "context"
    "errors"
    "fmt"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "gorm.io/gorm/utils"
    "log"
    "os"
    "time"
)

// GetDB 获取db实例
func GetDB(name string, c *gin.Context) *gorm.DB {
    db := DBMap[name]
    if ProjectConfig.Environment != EnvProd {
        debugLogger := New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
            SlowThreshold:             200 * time.Millisecond,
            LogLevel:                  logger.Info,
            IgnoreRecordNotFoundError: false,
            Colorful:                  true,
        })
        Log.Info("非PROD环境，日志级别设置为info")
        db := DBMap[name]
        if traceID, ok := c.Get(TraceKey); ok {
            db.Statement.Context = context.WithValue(context.Background(), TraceKey, traceID)
        }
        db.Logger = debugLogger
    }
    return db
}

// New initialize logger
func New(writer logger.Writer, config logger.Config) logger.Interface {
    var (
        infoStr      = "%s\n[info] "
        warnStr      = "%s\n[warn] "
        errStr       = "%s\n[error] "
        traceStr     = "%s\n[%.3fms] [rows:%v] %s [traceID:%s]"
        traceWarnStr = "%s %s\n[%.3fms] [rows:%v] %s [traceID:%s]"
        traceErrStr  = "%s %s\n[%.3fms] [rows:%v] %s [traceID:%s]"
    )

    if config.Colorful {
        infoStr = logger.Green + "%s\n" + logger.Reset + logger.Green + "[info] " + logger.Reset
        warnStr = logger.BlueBold + "%s\n" + logger.Reset + logger.Magenta + "[warn] " + logger.Reset
        errStr = logger.Magenta + "%s\n" + logger.Reset + logger.Red + "[error] " + logger.Reset
        traceStr = logger.Green + "%s\n" + logger.Reset + logger.Yellow + "[%.3fms] " +
            logger.BlueBold + "[rows:%v]" + logger.Reset + " %s" + logger.Green + " [traceID:%s]" + logger.Reset

        traceWarnStr = logger.Green + "%s " + logger.Yellow + "%s\n" + logger.Reset + logger.RedBold + "[%.3fms] " +
            logger.Yellow + "[rows:%v]" + logger.Magenta + " %s" + logger.Green + " [traceID:%s]" + logger.Reset

        traceErrStr = logger.RedBold + "%s " + logger.MagentaBold + "%s\n" + logger.Reset + logger.Yellow + "[%.3fms]" +
            logger.BlueBold + " [rows:%v]" + logger.Reset + " %s" + logger.Green + " [traceID:%s]" + logger.Reset
    }

    return &traceLogger{
        Writer:       writer,
        Config:       config,
        infoStr:      infoStr,
        warnStr:      warnStr,
        errStr:       errStr,
        traceStr:     traceStr,
        traceWarnStr: traceWarnStr,
        traceErrStr:  traceErrStr,
    }
}

type traceLogger struct {
    logger.Writer
    logger.Config
    infoStr, warnStr, errStr            string
    traceStr, traceErrStr, traceWarnStr string
}

// LogMode log mode
func (l *traceLogger) LogMode(level logger.LogLevel) logger.Interface {
    newlogger := *l
    newlogger.LogLevel = level
    return &newlogger
}

// Info print info
func (l traceLogger) Info(ctx context.Context, msg string, data ...interface{}) {
    if l.LogLevel >= logger.Info {
        l.Printf(l.infoStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
    }
}

// Warn print warn messages
func (l traceLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
    if l.LogLevel >= logger.Warn {
        l.Printf(l.warnStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
    }
}

// Error print error messages
func (l traceLogger) Error(ctx context.Context, msg string, data ...interface{}) {
    if l.LogLevel >= logger.Error {
        l.Printf(l.errStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
    }
}

// Trace print sql message
func (l traceLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
    if l.LogLevel <= logger.Silent {
        return
    }

    traceID := ""
    trace := ctx.Value(TraceKey)
    if trace != nil {
        traceID, _ = trace.(string)
    }

    elapsed := time.Since(begin)
    switch {
    case err != nil && l.LogLevel >= logger.Error && (!errors.Is(err, logger.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
        sql, rows := fc()
        if rows == -1 {
            l.Printf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql, traceID)
        } else {
            l.Printf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql, traceID)
        }
    case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
        sql, rows := fc()
        slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
        if rows == -1 {
            l.Printf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql, traceID)
        } else {
            l.Printf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql, traceID)
        }
    case l.LogLevel == logger.Info:
        sql, rows := fc()
        if rows == -1 {
            l.Printf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql, traceID)
        } else {
            l.Printf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql, traceID)
        }
    }
}
