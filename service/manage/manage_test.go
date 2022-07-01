package manage

import (
    "dibulido-srv/dto/common"
    "dibulido-srv/global"
    "fmt"
    "github.com/stretchr/testify/assert"
    "strconv"
    "testing"
    "time"
)

// TestCasbinService_Casbin 测试casbin实例生成
func TestCasbinService_CreatePolicy(t *testing.T) {
    casbinService := initEnv()
    casbinInfos := []*common.CasbinInfo{
        {
            Path:   "/v1/user/close",
            Method: "post",
        },
    }
    roleID := strconv.FormatInt(time.Now().Unix(), 10)
    err := casbinService.CreatePolicy(roleID, casbinInfos)
    assert.True(t, err == nil)
    casbinInfos = []*common.CasbinInfo{
        {
            Path:   "/v1/user/close",
            Method: "get",
        },
    }
    err = casbinService.CreatePolicy(roleID, casbinInfos)
    assert.True(t, err == nil)
    pathMaps := casbinService.GetPolicyPathByRoleID(roleID)
    fmt.Println(pathMaps)
}

// TestCasbinService_GetPolicyPathByRoleID 测试角色权限列表查询
func TestCasbinService_GetPolicyPathByRoleID(t *testing.T) {
    casbinService := initEnv()
    pathMaps := casbinService.GetPolicyPathByRoleID("guest")
    fmt.Println(pathMaps)
}

func initEnv() CasbinService {
    casbinService := CasbinService{}
    global.InitViper("../../conf/config.yaml")
    global.InitLog()
    global.InitDB()
    global.ProjectConfig.Casbin.ModelPath = "../../conf/rbac_model.conf"
    return casbinService
}

// TestCasbinService_AddRolesForUser
func TestCasbinService_AddRolesForUser(t *testing.T) {
    casbinService := initEnv()
    err := casbinService.AddRolesForUser("9e789586-93f3-406a-8d80-787d87cfe239", "root")
    assert.True(t, err == nil)

    info := &common.CasbinInfo{
        Path:   "/pms/grant",
        Method: "POST",
    }
    err = casbinService.CreatePolicy("root", []*common.CasbinInfo{info})
    assert.True(t, err == nil)
}
