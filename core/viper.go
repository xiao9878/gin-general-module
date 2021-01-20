package core

import (
	"fmt"
	"gin-general-module/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper(path string) *viper.Viper {
	var config string
	v := viper.New()
	v.SetConfigFile(config)
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
		fmt.Println(err)
	}
	return v
}
