package models

import (
	"golang.org/x/crypto/bcrypt"
	"test_viper/config"
)

func Login(userInfo UserInfo) (bool, string) {
	var isExist UserInfo
	config.DB.Where("user = ?", userInfo.User).First(&isExist)
	//判断是否存在
	if isExist.ID == 0 {
		return false, "账号不存在"
	}
	//判断账号密码是否正确
	err := bcrypt.CompareHashAndPassword([]byte(isExist.Password), []byte(userInfo.Password))
	if err != nil {
		return false, "密码错误"
	}
	return true, ""
}
