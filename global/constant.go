package global

const (
    ConfigEnv  = "dibu_config"      // 环境系统变量key
    ConfigFile = "conf/config.yaml" // 配置文件地址
)

// 环境枚举
const (
    EnvDev     = "DEV"
    EnvTest    = "TEST"
    EnvSandbox = "SANDBOX"
    EnvProd    = "PROD"
)

// 数据库名称列表
const (
    ConfigDB = "config"
    DataDB   = "data"
)

// GinUserInfoKey 用户信息在gin context中的key
const GinUserInfoKey = "userInfo"

// TraceKey gin框架中的traceID
const TraceKey = "traceID"

// Token 在请求header中的key
const (
    TokenHeaderKey  = "x-Authorization"
    TokenExpiredKey = "x-expired-at"
)
