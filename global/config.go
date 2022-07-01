// Package global 一些全局的设置比如 log configuration 和一些全局定义的变量和常量
package global

import (
    "dibulido-srv/dto/common"
)

// ProjectConfig 全局配置
var ProjectConfig Configuration

// Configuration 全局配置信息
type Configuration struct {
    Port         string                  `mapstructure:"post" yaml:"port" json:"port"`
    AppName      string                  `mapstructure:"app" json:"app" yaml:"app"`          // 应用名称
    Environment  string                  `mapstructure:"env" json:"env" yaml:"env"`          // 环境 DEV/TEST/SANDBOX/PROD
    LoggerConfig Options                 `mapstructure:"logger" yaml:"logger" json:"logger"` // 日志配置
    GormList     map[string]common.Mysql `mapstructure:"gorm" yaml:"gorm" json:"gorm"`       // 数据库链接配置
    Jwt          JwtConfig               `mapstructure:"jwt" yaml:"jwt" json:"jwt"`          // auth 配置信息
    Casbin       common.CasbinConfig     `mapstructure:"casbin" yaml:"casbin" json:"casbin"` // Casbin配置
}
