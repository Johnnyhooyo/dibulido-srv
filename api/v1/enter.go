package v1

import (
    "dibulido-srv/api/v1/backend"
    "dibulido-srv/api/v1/manage"
)

// ApiGroup 所有接口
type ApiGroup struct {
    BackendGroup backend.ApiGroup
    ManageGroup  manage.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
