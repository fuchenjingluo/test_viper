package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"test_viper/models"
	"test_viper/utils"
)

// 获取用户数据
func getParam(c *gin.Context, key string) string {
	if value := c.Query(key); value != "" {
		return value
	}
	return c.PostForm(key)
}
func Registered(c *gin.Context) {
	// 获取用户输入
	user := getParam(c, "user")
	password := getParam(c, "password")

	// 初始化用户信息
	var userInfo models.UserInfo
	userInfo.User = user
	userInfo.Password = password
	// 验证并绑定参数
	if err := utils.Validate.Struct(userInfo); err != nil {
		// 获取验证错误并进行翻译
		errors := err.(validator.ValidationErrors)
		translatedErrors := errors.Translate(utils.Trans)
		c.JSON(http.StatusOK, gin.H{
			"msg": translatedErrors,
		})
		return
	}
	//密码加密
	hashPassword, err := utils.HashPassword(password)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hashPassword)
	userInfo.Password = hashPassword
	//检查用户是否存在
	if !models.Register(userInfo) {
		c.JSON(http.StatusOK, gin.H{
			"mes": "账号已存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
}
