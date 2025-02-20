package initialize

import (
	"fmt"

	"github.com/Vzttfromxduszu/golang-1.git/common/global"
	"github.com/spf13/viper"
)

// LoadConfig 加载配置文件
func LoadConfig() {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	if err := viper.Unmarshal(&global.Config); err != nil { // 将配置文件的内容绑定到 global.Config 结构体
		panic(fmt.Errorf("unable to decode into struct %w", err)) // 如果绑定失败，抛出错误并终止程序
	}
}
