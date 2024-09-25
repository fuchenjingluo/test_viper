package api

import (
	"github.com/gin-gonic/gin"
	"test_viper/moddleware/jwt"
)

func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": jwt.Claims,
	})
}
