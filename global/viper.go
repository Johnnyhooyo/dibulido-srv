package global

import (
    "flag"
    "fmt"
    "github.com/fsnotify/fsnotify"
    "github.com/spf13/viper"
    "log"
    "os"
)

// Viper 全局配置
var Viper *viper.Viper

// InitViper 加载配置信息 优先级: 命令行 > 环境变量 > 默认值
func InitViper(path ...string) {
    var config string
    if len(path) == 0 {
        flag.StringVar(&config, "c", "", "choose config file.")
        flag.Parse()
        if config == "" {
            if configEnv := os.Getenv(ConfigEnv); configEnv == "" {
                config = ConfigFile
                fmt.Printf("您正在使用config的默认值,config的路径为%v\n", ConfigFile)
            } else {
                config = configEnv
                fmt.Printf("您正在使用GVA_CONFIG环境变量,config的路径为%v\n", config)
            }
        } else {
            fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
        }
    } else {
        config = path[0]
        fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", config)
    }

    // config = convertIfRelativePath(config)

    v := viper.New()
    v.SetConfigFile(config)
    v.SetConfigType("yaml")
    err := v.ReadInConfig()
    if err != nil {
        log.Fatalln("Fatal error config file:", err)
    }
    v.WatchConfig()

    v.OnConfigChange(func(e fsnotify.Event) {
        fmt.Println("config file changed:", e.Name)
        if err := v.Unmarshal(&ProjectConfig); err != nil {
            log.Fatalln("配置信息变更后拉取失败:", err.Error())
        }
    })
    if err := v.Unmarshal(&ProjectConfig); err != nil {
        log.Fatalln("配置信息拉取失败", err.Error())
    }
    Viper = v
}
