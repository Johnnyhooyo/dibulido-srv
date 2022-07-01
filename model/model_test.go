package model

import (
    "dibulido-srv/model/user"
    "github.com/magiconair/properties/assert"
    "testing"
)

// TestUser user测试
func TestUser(t *testing.T) {
    u := user.User{
        Username: "dibu",
        Password: "1111",
        Email:    "dibu@tencent.com",
    }
    userID, _ := u.Create()
    assert.Equal(t, userID, "")
}
