package middleware

import (
    response "dibulido-srv/api/model/repsonse"
    "dibulido-srv/global"
    errs "dibulido-srv/global/response"
    "dibulido-srv/service"
    "github.com/gin-gonic/gin"
)

var casbinService = service.IServiceGroupApp.ManageServiceGroup.CasbinService

// CasbinMiddleware 拦截器 需要在登陆中间件 jwt之后使用，不然获取不到用户信息
// 获取不到信息就检查接口资源是否存在policy存在就拦截
func CasbinMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 获取请求的PATH
        obj := c.Request.URL.Path
        // 获取请求方法
        act := c.Request.Method

        var userInfo *global.UserInfo
        if v, ok := c.Get(global.GinUserInfoKey); !ok {
            if policies := casbinService.Casbin().GetFilteredPolicy(1, obj); len(policies) > 0 {
                global.Log.Errorf("未登陆用户请求受限资源:%s", obj)
                response.FailWithError(errs.TokenInvalid, c)
                c.Abort()
                return
            } else {
                c.Next()
                return
            }
        } else {
            if userInfo, ok = v.(*global.UserInfo); !ok {
                global.Log.Errorf("非法token，未提取到用户信息:%s", obj)
                response.Fail(c)
                c.Abort()
                return
            }
        }

        // 获取用户的角色
        sub := userInfo.UserID

        // 判断策略中是否存在
        hasPermission, err := hasImplicitPermissionsForUser(sub, obj, act)
        if err != nil {
            global.Log.Errorf("查询权限失败：%s", err.Error())
            response.Fail(c)
        } else if hasPermission {
            c.Next()
        } else {
            response.FailWithError(errs.PermissionForbid, c)
            c.Abort()
        }
    }
}

func hasImplicitPermissionsForUser(sub, obj, act string) (bool, error) {
    e := casbinService.Casbin()
    roles, err := e.GetImplicitPermissionsForUser(sub)
    if err != nil {
        return false, err
    }
    for _, role := range roles {
        if len(role) < 3 {
            return false, errs.SystemError
        }
        if role[1] == obj && role[2] == act {
            return true, nil
        }
    }
    return false, nil
}
