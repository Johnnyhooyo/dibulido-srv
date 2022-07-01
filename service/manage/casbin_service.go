package manage

import (
    "dibulido-srv/global"
    "github.com/casbin/casbin/v2"
    gormadapter "github.com/casbin/gorm-adapter/v3"
    "sync"
)

type CasbinService struct{}

var CasbinServiceApp = new(CasbinService)

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
