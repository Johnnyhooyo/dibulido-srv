package service

import "dibulido-srv/service/manage"

// IServiceGroup 处理层
type IServiceGroup struct {
    ManageServiceGroup manage.ServiceGroup
}

// IServiceGroupApp 处理方法指针
var IServiceGroupApp = new(IServiceGroup)
