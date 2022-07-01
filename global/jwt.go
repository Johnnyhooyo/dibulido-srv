package global

import (
    rand2 "crypto/rand"
    "crypto/rsa"
    "dibulido-srv/global/response"
    "github.com/golang-jwt/jwt/v4"
    "io/ioutil"
    logger "log"
    "time"
)

var privateKey *rsa.PrivateKey

// JwtConfig jwt配置
type JwtConfig struct {
    ExpiredMinute  int64  `mapstructure:"expired_minute" json:"expiredMinute" yaml:"expired_minute"`
    PrivateKeyPath string `mapstructure:"private_path" json:"privatePath" yaml:"private_path"`
    PublicKeyPath  string `mapstructure:"public_path" json:"publicPath" yaml:"public_path"` // 如果可以从私钥中获取公钥 这个就可以不配置
    SingKey        string `mapstructure:"singKey" json:"singKey" yaml:"singKey"`            // 现在使用RSA签名，如果用HMAC则用得到
    sinKey         *rsa.PrivateKey
}

// UserInfo 用户信息
type UserInfo struct {
    UserID   string
    UserName string
}

// JwtClaims token内容
type JwtClaims struct {
    *jwt.RegisteredClaims
    TokenType string
    *UserInfo
}

// InitJwtFromConf 读取配置文件中的密钥
func InitJwtFromConf(path ...string) {
    var err error
    var pk []byte
    if len(path) > 0 {
        pk, err = ioutil.ReadFile(path[0])
    } else {
        pk, err = ioutil.ReadFile(ProjectConfig.Jwt.PrivateKeyPath)
    }
    if err != nil {
        logger.Fatalln("获取私钥错误:", err.Error())
    }
    privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(pk)
    if err != nil {
        logger.Fatalln("获取私钥错误:", err.Error())
    }
}

// InitJwt 初始化jwt需要的密钥
// 这里可以生成密钥，但是是单机的版本
func InitJwt() {
    pk, err := rsa.GenerateKey(rand2.Reader, 1024)
    if err != nil {
        logger.Fatalln("获取私钥错误:", err.Error())
    } else {
        privateKey = pk
    }
}

// Sign jwt生成token
func Sign(u *UserInfo) (token string, err error) {
    t := jwt.New(jwt.GetSigningMethod("RS256"))
    t.Claims = &JwtClaims{
        &jwt.RegisteredClaims{
            Issuer:    ProjectConfig.AppName,
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(ProjectConfig.Jwt.ExpiredMinute))),
        },
        "base",
        u,
    }
    token, err = t.SignedString(privateKey)
    if err != nil {
        return "", err
    }
    return token, nil
}

// VerifySign 验签
func VerifySign(tokenStr string) (*JwtClaims, error) {
    token, err := jwt.ParseWithClaims(tokenStr, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
        return &privateKey.PublicKey, nil
    })
    if err != nil {
        if ve, ok := err.(*jwt.ValidationError); ok {
            if ve.Errors&jwt.ValidationErrorExpired != 0 {
                // Token is expired
                return tokenToClaims(token)
            } else {
                return nil, response.TokenInvalid
            }
        }
    }
    return tokenToClaims(token)
}

func tokenToClaims(token *jwt.Token) (*JwtClaims, error) {
    if token != nil {
        if claims, ok := token.Claims.(*JwtClaims); ok {
            if !token.Valid {
                return claims, response.TokenExpired
            }
            return claims, nil
        }
        return nil, response.TokenInvalid
    } else {
        return nil, response.TokenInvalid
    }
}
