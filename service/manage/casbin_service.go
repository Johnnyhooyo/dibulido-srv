package manage

import (
    "dibulido-srv/dto/common"
    "dibulido-srv/global"
    "dibulido-srv/global/response"
    "github.com/casbin/casbin/v2"
    gormadapter "github.com/casbin/gorm-adapter/v3"
    "sync"
)

// CasbinService casbin服务
type CasbinService struct{}

// 角色管理

// CreatePolicy 创建权限策略 角色存的话就是追加权限
func (casbinService *CasbinService) CreatePolicy(roleID string, casbinInfos []*common.CasbinInfo) error {
    var rules [][]string
    for _, v := range casbinInfos {
        rules = append(rules, []string{roleID, v.Path, v.Method})
    }
    e := casbinService.Casbin()
    if _, err := e.AddPolicies(rules); err != nil {
        return response.CasbinError
    }
    return nil
}

// UpdatePolicy 全量替换角色权限策略
func (casbinService *CasbinService) UpdatePolicy(roleID string, casbinInfos []*common.CasbinInfo) error {
    casbinService.ClearCasbin(0, roleID)
    return casbinService.CreatePolicy(roleID, casbinInfos)
}

// RemovePolicy 删除角色指定权限策略
func (casbinService *CasbinService) RemovePolicy(roleID string, casbinInfos []*common.CasbinInfo) error {
    e := casbinService.Casbin()
    var rules [][]string
    for _, v := range casbinInfos {
        rules = append(rules, []string{roleID, v.Path, v.Method})
    }
    if ok, err := e.RemovePolicy(rules); err != nil || !ok {
        return response.CasbinError
    }
    return nil
}

// GetPolicyPathByRoleID 获取对应角色的权限列表
func (casbinService *CasbinService) GetPolicyPathByRoleID(roleID string) (pathMaps []*common.CasbinInfo) {
    e := casbinService.Casbin()
    list := e.GetFilteredPolicy(0, roleID)
    for _, v := range list {
        pathMaps = append(pathMaps, &common.CasbinInfo{
            Path:   v[1],
            Method: v[2],
        })
    }
    return pathMaps
}

// ClearCasbin 清楚对应角色的权限信息
func (casbinService *CasbinService) ClearCasbin(fieldIndex int, fieldValues ...string) bool {
    e := casbinService.Casbin()
    success, _ := e.RemoveFilteredPolicy(fieldIndex, fieldValues...)
    return success
}

// 用户角色绑定 多对多

// AddRolesForUser 给用户绑定角色
func (casbinService *CasbinService) AddRolesForUser(userID string, roleIDs ...string) error {
    e := casbinService.Casbin()
    _, err := e.AddRolesForUser(userID, roleIDs)
    if err != nil {
        return response.CasbinError
    }
    return nil
}

// GetRolesForUser 获取用户所有角色信息
func (casbinService *CasbinService) GetRolesForUser(userID string) (*common.UserRoleInfo, error) {
    e := casbinService.Casbin()
    if roles, err := e.GetRolesForUser(userID); err != nil {
        return nil, response.CasbinError
    } else {
        return &common.UserRoleInfo{
            UserID:  userID,
            RoleIDs: roles,
        }, nil
    }
}

// DeleteRoleForUser 删除用户指定角色
func (casbinService *CasbinService) DeleteRoleForUser(userID string, roleID string) error {
    e := casbinService.Casbin()
    _, err := e.DeleteRoleForUser(userID, roleID)
    if err != nil {
        return response.CasbinError
    }
    return nil
}

var (
    syncedEnforcer *casbin.SyncedEnforcer
    once           sync.Once
)

// Casbin 获取casbin实例
func (casbinService *CasbinService) Casbin() *casbin.SyncedEnforcer {
    once.Do(func() {
        a, _ := gormadapter.NewAdapterByDB(global.DBMap[global.ConfigDB])
        syncedEnforcer, _ = casbin.NewSyncedEnforcer(global.ProjectConfig.Casbin.ModelPath, a)
    })
    _ = syncedEnforcer.LoadPolicy()
    return syncedEnforcer
}
