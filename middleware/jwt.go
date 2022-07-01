package middleware

import (
    response "dibulido-srv/api/model/repsonse"
    "dibulido-srv/global"
    "dibulido-srv/global/log"
    errs "dibulido-srv/global/response"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "strconv"
    "time"
)

const userIDKey = "userID"

// JwtMiddleware jwt登陆检查
func JwtMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 我们这里jwt鉴权取头部信息 x-Authorization 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
        token := c.Request.Header.Get(global.TokenHeaderKey)
        if token == "" {
            response.FailWithError(errs.TokenInvalid, c)
            c.Abort()
            return
        }

        // parseToken 解析token包含的信息
        claims, err := global.VerifySign(token)
        if err != nil {
            if err == errs.TokenExpired {
                if claims.ExpiresAt.Unix()-time.Now().Unix() < 60*10 {
                    if err := refreshToken(claims, c); err != nil {
                        response.FailWithError(err, c)
                        c.Abort()
                        return
                    }
                }

            } else {
                response.FailWithError(errs.TokenInvalid, c)
                c.Abort()
                return
            }
        }
        c.Set(global.GinUserInfoKey, claims.UserInfo)
        // 日志追加userID
        log.NewContext(c, zap.String(userIDKey, claims.UserID))
        c.Next()
    }
}

// refreshToken 给c一个新的token
func refreshToken(claims *global.JwtClaims, c *gin.Context) error {
    token, err := global.Sign(claims.UserInfo)
    if err != nil {
        return err
    } else {
        newClaims, err := global.VerifySign(token)
        if err != nil {
            return err
        }
        c.Header(global.TokenHeaderKey, token)
        c.Header(global.TokenExpiredKey, strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
    }
    return nil
}
