package util

import (
    response "dibulido-srv/api/model/repsonse"
    "dibulido-srv/global/log"
    "github.com/gin-gonic/gin"
)

// BindParam 绑定请求中的参数到obj 绑定失败返回false
func BindParam(obj interface{}, c *gin.Context) bool {
    if err := c.ShouldBindJSON(obj); err != nil {
        log.WithContext(c).Errorf("获取请求参数失败,%s", err.Error())
        response.ParamCheckFail(err, c)
        return false
    }
    return true
}
