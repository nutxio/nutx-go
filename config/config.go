package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var VP *viper.Viper

func init() {
	v := viper.New()
	v.AddConfigPath("config")
	v.SetConfigName("app")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("配置文件读取失败：%s", err))
	}

	VP = v
}
