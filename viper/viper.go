package viper

import (
	"fmt"
	"github.com/spf13/viper"
	"utils/global"
)

func Viper() {
	viper.SetConfigFile("./config.dev.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	viper.Unmarshal(&global.Config)
	fmt.Println("读取viper配置", global.Config)
}
