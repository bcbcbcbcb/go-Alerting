package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

var Config *viper.Viper

func config_init() {
	v := viper.New()
	Config = v

	Config.SetConfigName("conf")     // 配置文件名称，不需要后缀名
	Config.SetConfigType("yaml")     // 配置文件类型；支持 JSON, TOML, YAML, HCL, INI, envfile or Java properties formats
	Config.AddConfigPath("./config") // 配置文件查找路径1
	Config.AddConfigPath(".")        // 配置文件查找路径2

	if err := Config.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// 配置默认值
	Config.SetDefault("app.address", "0.0.0.0")
	Config.SetDefault("app.port", 8888)
	Config.SetDefault("http.proxy", "nil")

	// fmt.Println(Config.GetString("http.proxy"))
	// fmt.Println(Config.GetString("app.address"))
	// fmt.Println(Config.GetString("app.port"))
}
