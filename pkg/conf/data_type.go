package conf

// 系统配置
type system struct {
	Port          string
	Debug         bool
	SessionSecret string
	HashIDSalt    string
}

// 数据库配置
type database struct {
	Type        string
	Host        string
	Port        uint
	User        string
	Password    string
	Database    string
	Charset     string
	File        string
	TablePrefix string
}

// 跨域配置
type cors struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	AllowCredentials bool
	ExposeHeaders    []string
}
