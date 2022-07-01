// Package manage 服务实现的地方，如果有不同的实现方法，可以定义一个接口，然后有不同的实现方法。
// 在获取时候根据配置或者参数去决定使用哪一个实现
// 如casbin 是实现了权限功能 定义 type permission interface{} 然后在new时候选取casbin
package manage

// ServiceGroup 管理方法集合
type ServiceGroup struct {
    UserService
    LogService
    CasbinService
}
