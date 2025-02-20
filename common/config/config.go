package config

// 组合全部配置模型
type Config struct {
	Server Server `mapstructure:"server" json:"server" yaml:"server"`
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}

type Server struct {
	Post string `mapstructure:"post" json:"post" yaml:"post"`
}

type Mysql struct {
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Url      string `mapstructure:"url" json:"url" yaml:"url"`
}
