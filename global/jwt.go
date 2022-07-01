package auth

import (
    rand2 "crypto/rand"
    "crypto/rsa"
    "dibulido-srv/global"
    "dibulido-srv/global/log"
    "dibulido-srv/global/response"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
    logger "log"
    "time"
)

var privateKey *rsa.PrivateKey

// JwtConfig jwt配置
type JwtConfig struct {
    SingKey       string `mapstructure:"singKey" json:"singKey" yaml:"singKey"` // 现在使用RSA签名，如果用HMAC则用得到
    ExpiredMinute int64  `mapstructure:"expiredMinute" json:"expiredMinute" yaml:"expiredMinute"`
    sinKey        *rsa.PrivateKey
}

// UserInfo 用户信息
type UserInfo struct {
    UserID   string
    UserName string
    // todo 权限信息
}

// JwtClaims token内容
type JwtClaims struct {
    *jwt.RegisteredClaims
    TokenType string
    *UserInfo
}

// InitJwt 初始化jwt需要的密钥
func InitJwt() {
    pk, err := rsa.GenerateKey(rand2.Reader, 1024)
    if err != nil {
        logger.Fatalln("获取私钥错误:", err.Error())
    } else {
        privateKey = pk
    }
}

// Sign jwt生成token
func Sign(u *UserInfo, c *gin.Context) (token string, err error) {
    t := jwt.New(jwt.GetSigningMethod("RS256"))
    t.Claims = &JwtClaims{
        &jwt.RegisteredClaims{
            Issuer:    global.ProjectConfig.AppName,
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(global.ProjectConfig.Jwt.ExpiredMinute))),
        },
        "base",
        u,
    }
    token, err = t.SignedString(privateKey)
    if err != nil {
        log.WithContext(c).Errorf("签名错误:%s", err.Error())
        return "", err
    }
    return token, nil
}

// Check 验签
func VierfySgin(tokenStr string, c *gin.Context) (*JwtClaims, error) {
    token, err := jwt.ParseWithClaims(tokenStr, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
        return privateKey.PublicKey, nil
    })
    if err != nil {
        if ve, ok := err.(*jwt.ValidationError); ok {
            if ve.Errors&jwt.ValidationErrorExpired != 0 {
                // Token is expired
                return nil, response.TokenExpired
            } else {
                return nil, response.TokenInvalid
            }
        }
    }
    if token != nil {
        if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
            return claims, nil
        }
        return nil, response.TokenInvalid
    } else {
        return nil, response.TokenInvalid
    }
}
