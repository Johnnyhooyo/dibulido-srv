# dibulido-srv
❄️ 寒冬之中仍有雪和冰的美丽，保持虔诚，以待春之水暖。 

这里是一个Golang的脚手架项目，帮助快速开发一个Web项目。

## 技术栈列表
1. [GIN](https://github.com/gin-gonic/gin#dont-trust-all-proxies) - 一个高性能的网络框架 `gin-gonic/gin: Gin is a HTTP web framework written in Go (Golang). It features a Martini-like API with much better performance -- up to 40 times faster. If you need smashing performance, get yourself some Gin.`
2. [swagger](https://github.com/go-swagger/go-swagger) 文档生成 `Swagger is a simple yet powerful representation of your RESTful API.`
3. [zap](https://github.com/uber-go/zap) 日志框架 `Blazing fast, structured, leveled logging in Go.`
4. [lumberjack](https://github.com/natefinch/lumberjack) 按大小切分的滚动日志框架 `Lumberjack is a Go package for writing logs to rolling files.`
5. [viper](https://github.com/spf13/viper) 配置加载工具，可读取配置文件、系统变量和远程配置信息；支持热更新 `Go configuration with fangs!`
6. [casbin](https://github.com/casbin/casbin) 权限管理工具 `Casbin is a powerful and efficient open-source access control library for Golang projects. It provides support for enforcing authorization based on various access control models.`
7. [gorm](https://github.com/go-gorm/gorm) 数据库链接管理工具 `The fantastic ORM library for Golang, aims to be developer friendly.`
8. [jwt](https://github.com/golang-jwt/jwt) 验证工具 `A go (or 'golang' for search engine friendliness) implementation of JSON Web Tokens.`
9. [govalidator](https://github.com/asaskevich/govalidator) 参数校验(未使用) `A package of validators and sanitizers for strings, structs and collections. Based on validator.js.`
10. [validator](https://github.com/go-playground/validator) [文档](https://github.com/go-playground/validator/blob/dd2857a4cb6c53af5bf6944cc15fd04c865d28ae/doc.go) 参数校验(gin中默认有此组件 对应版本v10.6.1) `Package validator implements value validations for structs and individual fields based on tags.`
11. [validate](https://github.com/go-openapi/validate) 辅助swagger2.0的参数校验？ `This package provides helpers to validate Swagger 2.0. specification (aka OpenAPI 2.0).`

## License
[MIT License](LICENSE)