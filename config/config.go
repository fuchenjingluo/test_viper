package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string
		Port string
	}
	Database struct {
		Dsn          string
		MaxOpenConns int
		MaxIdleConns int
	}
}

var AppConfig *Config

func InitConfig() {
	viper.SetConfigName("config")    // 配置文件名称(无扩展名)
	viper.SetConfigType("yml")       // 扩展名
	viper.AddConfigPath("./config/") // 指定配置文件路径
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	AppConfig = &Config{}
	err = viper.Unmarshal(AppConfig) // 将配置解析到 AppConfig 结构体中
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	InitDB()
	InitRedis()
}
