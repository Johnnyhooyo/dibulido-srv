package common

// CasbinConfig casbin配置文件
type CasbinConfig struct {
    ModelPath string `mapstructure:"model_path" json:"modelPath" yaml:"model_path"` // 存放casbin模型的相对路径
}

// CasbinInfo Casbin info structure
type CasbinInfo struct {
    Path   string `json:"path"`   // 路径
    Method string `json:"method"` // 方法
}

// CasbinInReceive Casbin structure for input parameters
type CasbinInReceive struct {
    RoleID      string       `json:"roleID"` // 权限ID
    CasbinInfos []CasbinInfo `json:"casbinInfos"`
}

// UserRoleInfo 用户角色关系 1对多
// swagger:model userRole
type UserRoleInfo struct {
    UserID string `json:"userID"  binding:"required"`
    // min size: 1
    RoleIDs []string `json:"roleIDs" binding:"gt=0,required"`
}
