package models

import (
	"gorm.io/gorm"
	"test_viper/config"
)

type UserInfo struct {
	gorm.Model
	User     string `json:"user" gorm:"unique" validate:"required,min=4,max=32"`
	Password string `json:"password" validate:"required,min=6,max=18"`
}

func Register(userInfo UserInfo) bool {
	config.DB.AutoMigrate(&UserInfo{})
	//判断账号是否存在
	var isExist UserInfo
	config.DB.Where("user = ?", userInfo.User).First(&isExist)
	if isExist.ID > 0 {
		return false
	}
	config.DB.Create(&userInfo)
	return true
}
