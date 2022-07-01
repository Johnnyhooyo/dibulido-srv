package manage

import "dibulido-srv/service"

// ApiGroup 管理接口集合
type ApiGroup struct {
    IUserApi
    LogApi
    PermissionAPI
}

var (
    userService   = service.IServiceGroupApp.ManageServiceGroup.UserService
    logService    = service.IServiceGroupApp.ManageServiceGroup.LogService
    casbinService = service.IServiceGroupApp.ManageServiceGroup.CasbinService
)
