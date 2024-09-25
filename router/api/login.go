package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"test_viper/models"
	"test_viper/utils"
)

func Login(c *gin.Context) {
	var login models.UserInfo
	login.User = getParam(c, "user")
	login.Password = getParam(c, "password")
	err := utils.Validate.Struct(login)
	// 使用翻译后的错误信息
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.(validator.ValidationErrors).Translate(utils.Trans)})
		return
	}
	if exist, errors := models.Login(login); !exist {
		c.JSON(http.StatusOK, gin.H{
			"msg": errors,
		})
		return
	}
	token, err := utils.GenerateJwt(login.User)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"msg":   "登陆成功",
		"token": token,
	})
}
