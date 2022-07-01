package global

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "gorm.io/plugin/dbresolver"
    "log"
)

// DBMap 数据库连接实例map
var DBMap = make(map[string]*gorm.DB)

// DBConnection 数据库连接代理
type DBConnection struct {
    err    error
    baseDB *gorm.DB
}

func (c DBConnection) checkError() {
    if c.err != nil {
        log.Fatalln("DB init err:", c.err.Error())
    }
}

// InitDB 初始化所有DB链接
func InitDB() {
    conn := DBConnection{}
    debugLogger := logger.Default
    if ProjectConfig.Environment != EnvProd {
        debugLogger = logger.Default.LogMode(logger.Info)
        Log.Info("非PROD环境，日志级别设置为info")
    }
    for name, gormConfig := range ProjectConfig.GormList {
        conn.baseDB, conn.err = gorm.Open(mysql.New(mysql.Config{
            DSN: gormConfig.Dsn(),
        }), &gorm.Config{
            Logger: debugLogger,
        })
        conn.checkError()
        conn.err = conn.baseDB.Use(
            dbresolver.Register(dbresolver.Config{}).
                SetMaxIdleConns(gormConfig.MaxIdleConn).
                SetMaxOpenConns(gormConfig.MaxOpenConn),
        )
        conn.checkError()
        DBMap[name] = conn.baseDB
    }

}
