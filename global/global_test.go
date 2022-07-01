package global

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

// TestInitJwt 验证jwt
func TestInitJwt(t *testing.T) {
    InitViper("../conf/config.yaml")
    InitJwtFromConf("../conf/private.key")
    assert.True(t, true)

    u := &UserInfo{
        UserID:   "111",
        UserName: "dfsdf",
    }
    sign, err := Sign(u)
    assert.True(t, err == nil)
    verifySign, err := VerifySign(sign)
    assert.True(t, err == nil)
    fmt.Println(verifySign.UserInfo)
}
