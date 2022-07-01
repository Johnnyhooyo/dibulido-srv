package router

import (
    "dibulido-srv/api/router/backend"
    "dibulido-srv/api/router/manage"
)

// BaseRouterGroup 基础路由集合
type BaseRouterGroup struct {
    Backend backend.RouterGroup
    Manage  manage.RouterGroup
}

var BaseRouterGroupApp = new(BaseRouterGroup)
