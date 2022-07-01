package main

import (
    "dibulido-srv/api/core"
    "dibulido-srv/global"
    "net/http"
    "time"
)

// @title dibulido_srv API
// @version 1.0
// @description This is a learning serve
// @termsOfService dibulido personally

// @contact.name dibulido
// @contact.email 359332997@qq.com

// @license.name MIT License
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey ApiKeyAuth
// @name x-Authorization
// @in header

// @host localhost:8080
// @BasePath /
func main() {
    global.InitViper()
    global.InitLog()
    global.InitDB()
    global.InitJwtFromConf()

    r := core.InitGinRouter()
    s := &http.Server{
        Addr:           ":8080",
        Handler:        r,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    global.Log.Info("server start success~")
    s.ListenAndServe()
}
