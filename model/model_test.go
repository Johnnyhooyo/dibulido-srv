package model

import (
    "dibulido-srv/global"
    "dibulido-srv/model/manage"
    "github.com/gofrs/uuid"
    "github.com/stretchr/testify/assert"
    "testing"
)

// TestUser user测试
func TestUser(t *testing.T) {
    global.InitViper("../conf/config.yaml")
    global.InitDB()
    global.InitLog()
    uuidV4, _ := uuid.NewV4()
    u := &manage.User{
        UserID:   uuidV4.String(),
        UserName: "dibu",
        Password: "1111",
        Email:    "dibu@tencent.com",
    }
    err := u.Create(nil)
    assert.True(t, err == nil)

}
