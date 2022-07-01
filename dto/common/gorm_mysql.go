package common

// Mysql mysql数据库连接信息
type Mysql struct {
    Host        string `mapstructure:"host" json:"host" yaml:"host"`                            // 服务器地址
    Port        string `mapstructure:"port" json:"port" yaml:"port"`                            // 端口
    Config      string `mapstructure:"config" json:"config" yaml:"config"`                      // 高级配置
    Dbname      string `mapstructure:"db_name" json:"db_name" yaml:"db_name"`                   // 数据库名
    Username    string `mapstructure:"username" json:"username" yaml:"username"`                // 数据库用户名
    Password    string `mapstructure:"password" json:"password" yaml:"password"`                // 数据库密码
    MaxIdleConn int    `mapstructure:"max_idle_conn" json:"max_idle_conn" yaml:"max_idle_conn"` // 空闲中的最大连接数
    MaxOpenConn int    `mapstructure:"max_open_conn" json:"max_open_conn" yaml:"max_open_conn"` // 打开到数据库的最大连接数
}

// Dsn 获取数据库链接url
func (m *Mysql) Dsn() string {
    return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}
