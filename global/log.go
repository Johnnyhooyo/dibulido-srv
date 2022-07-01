package log

import (
    "dibulido-srv/global"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "os"
    "path/filepath"
    "sync"
    "time"

    lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

// Options 日志配置参数
type Options struct {
    LogFileDir    string `json:"logFileDir" yaml:"logFileDir"` // 日志文件路径
    AppName       string `json:"appName" yaml:"appName"`       // 日志文件前缀
    ErrorFileName string `json:"errorFileName" yaml:"errorFileName"`
    WarnFileName  string `json:"warnFileName" yaml:"warnFileName"`
    InfoFileName  string `json:"infoFileName" yaml:"infoFileName"`
    DebugFileName string `json:"debugFileName" yaml:"debugFileName"`
    MaxSize       int    `json:"maxSize" yaml:"maxSize"`       // 单个文件最大存储量 MB
    MaxAge        int    `json:"maxAge" yaml:"maxAge"`         // 文件最多保存多少天
    MaxBackups    int    `json:"maxBackups" yaml:"maxBackups"` // 文件最多爆粗那个书
    Compress      bool   `json:"compress" yaml:"compress"`     // 文件是否压缩
    zap.Config
}

var (
    Log                            *Logger
    sp                             = string(filepath.Separator)
    errWS, warnWS, infoWS, debugWS zapcore.WriteSyncer
    debugConsoleWS                 = zapcore.Lock(os.Stdout)
    errorConsoleWS                 = zapcore.Lock(os.Stderr)
)

// InitLog 初始化日志
// Zap 提供了两种类型的日志记录器—Sugared Logger和Logger。
//
// 在性能很好但不是很关键的上下文中，使用SugaredLogger。它比其他结构化日志记录包快 4-10 倍，并且支持结构化和 printf 风格的日志记录。
// 在每一微秒和每一次内存分配都很重要的上下文中，使用Logger。它甚至比SugaredLogger更快，内存分配次数也更少，但它只支持强类型的结构化日志记录。
func InitLog(conf ...*Options) {
    Log = &Logger{
        Opts: &global.Config.LoggerConfig,
    }

    Log.Lock()
    defer Log.Unlock()
    if Log.initialized {
        Log.Info("[InitLog] Log initialized")
        return
    }

    if len(conf) > 0 {
        Log.Opts = conf[0]
    }
    Log.loadCfg()
    Log.init()
    Log.Info("[initLogger] zap plugin initializing completed")
    Log.initialized = true
}

// Logger 日志封装
type Logger struct {
    *zap.SugaredLogger
    sync.RWMutex
    initialized bool
    Opts        *Options `json:"opts"`
    zapConfig   zap.Config
}

func (l *Logger) init() {
    l.setSyncers()
    var err error
    myLogger, err := l.zapConfig.Build(l.cores())
    if err != nil {
        panic(err)
    }
    l.SugaredLogger = myLogger.Sugar()
    defer l.SugaredLogger.Sync()
}

func (l *Logger) loadCfg() {
    if l.Opts.Development {
        l.zapConfig = zap.NewDevelopmentConfig()
        l.zapConfig.EncoderConfig.EncodeTime = timeEncoder
    } else {
        l.zapConfig = zap.NewProductionConfig()
        l.zapConfig.EncoderConfig.EncodeTime = timeEncoder
    }
    if l.Opts.OutputPaths == nil || len(l.Opts.OutputPaths) == 0 {
        l.zapConfig.OutputPaths = []string{"stdout"}
    }
    if l.Opts.ErrorOutputPaths == nil || len(l.Opts.ErrorOutputPaths) == 0 {
        l.zapConfig.OutputPaths = []string{"stderr"}
    }
    // 默认输出到程序运行目录的logs子目录
    if l.Opts.LogFileDir == "" {
        l.Opts.LogFileDir, _ = filepath.Abs(filepath.Dir(filepath.Join(".")))
        l.Opts.LogFileDir += sp + "logs" + sp
    }
    if l.Opts.AppName == "" {
        l.Opts.AppName = "app"
    }
    if l.Opts.ErrorFileName == "" {
        l.Opts.ErrorFileName = "error.log"
    }
    if l.Opts.WarnFileName == "" {
        l.Opts.WarnFileName = "warn.log"
    }
    if l.Opts.InfoFileName == "" {
        l.Opts.InfoFileName = "info.log"
    }
    if l.Opts.DebugFileName == "" {
        l.Opts.DebugFileName = "debug.log"
    }
    if l.Opts.MaxSize == 0 {
        l.Opts.MaxSize = 100
    }
    if l.Opts.MaxBackups == 0 {
        l.Opts.MaxBackups = 30
    }
    if l.Opts.MaxAge == 0 {
        l.Opts.MaxAge = 5
    }
}

// lumberjack 按日志大小切割
// rotateLogs 按时间切割
func (l *Logger) setSyncers() {
    f := func(fN string) zapcore.WriteSyncer {
        return zapcore.AddSync(&lumberjack.Logger{
            Filename:   l.Opts.LogFileDir + sp + l.Opts.AppName + "-" + fN,
            MaxSize:    l.Opts.MaxSize,
            MaxBackups: l.Opts.MaxBackups,
            MaxAge:     l.Opts.MaxAge,
            Compress:   l.Opts.Compress,
            LocalTime:  true, // false 文件名的时间会是+0时区
        })
        // 每小时一个文件
        // logf, _ := rotateLogs.New(l.Opts.LogFileDir+sp+l.Opts.AppName+"-%Y%m%d_%H"+"-"+fN,
        //     rotateLogs.WithLinkName(l.Opts.LogFileDir+sp+l.Opts.AppName+"-"+fN),
        //     rotateLogs.WithMaxAge(30*24*time.Hour),
        //     rotateLogs.WithRotationTime(time.Minute),
        // )
        // return zapcore.AddSync(logf)
    }
    errWS = f(l.Opts.ErrorFileName)
    warnWS = f(l.Opts.WarnFileName)
    infoWS = f(l.Opts.InfoFileName)
    debugWS = f(l.Opts.DebugFileName)
    return
}

func (l *Logger) cores() zap.Option {
    fileEncoder := zapcore.NewJSONEncoder(l.zapConfig.EncoderConfig)
    // consoleEncoder := zapcore.NewConsoleEncoder(logger.zapConfig.EncoderConfig)
    encoderConfig := zap.NewDevelopmentEncoderConfig()
    encoderConfig.EncodeTime = timeEncoder
    consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

    errPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
        return lvl > zapcore.WarnLevel && zapcore.WarnLevel-l.zapConfig.Level.Level() > -1
    })
    warnPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
        return lvl == zapcore.WarnLevel && zapcore.WarnLevel-l.zapConfig.Level.Level() > -1
    })
    infoPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
        return lvl == zapcore.InfoLevel && zapcore.InfoLevel-l.zapConfig.Level.Level() > -1
    })
    debugPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
        return lvl == zapcore.DebugLevel && zapcore.DebugLevel-l.zapConfig.Level.Level() > -1
    })
    field := zap.String("ip", global.ServerIP()) // 添加日志发生机器IP
    cores := []zapcore.Core{
        zapcore.NewCore(fileEncoder, errWS, errPriority).With([]zap.Field{field}),
        zapcore.NewCore(fileEncoder, warnWS, warnPriority).With([]zap.Field{field}),
        zapcore.NewCore(fileEncoder, infoWS, infoPriority).With([]zap.Field{field}),
        zapcore.NewCore(fileEncoder, debugWS, debugPriority).With([]zap.Field{field}),
    }
    if l.Opts.Development {
        cores = append(cores, []zapcore.Core{
            zapcore.NewCore(consoleEncoder, errorConsoleWS, errPriority),
            zapcore.NewCore(consoleEncoder, debugConsoleWS, warnPriority),
            zapcore.NewCore(consoleEncoder, debugConsoleWS, infoPriority),
            zapcore.NewCore(consoleEncoder, debugConsoleWS, debugPriority),
        }...)
    }
    return zap.WrapCore(func(c zapcore.Core) zapcore.Core {
        return zapcore.NewTee(cores...)
    })
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
    enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func timeUnixNano(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
    enc.AppendInt64(t.UnixNano() / 1e6)
}
