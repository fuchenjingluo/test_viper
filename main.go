package main

import (
	"test_viper/config"
	"test_viper/router"
)

func main() {
	config.InitConfig()
	r := router.SetupRouter()
	r.Run(config.AppConfig.App.Port)
}
