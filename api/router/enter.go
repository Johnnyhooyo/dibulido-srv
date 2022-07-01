package v1

import "dibulido-srv/api/router/beckend"

type RouterGroup struct {
    Backend beckend.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
