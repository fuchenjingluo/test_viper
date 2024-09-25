package jwt

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"test_viper/utils"
)

var Claims interface{}

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.PostForm("token")
		if c.PostForm("token") == "" {
			token = c.Query("token")
		}
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"msg": "token不能为空",
			})
			c.Abort()
			return
		}
		claims, err := utils.ParseJwt(token)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			c.Abort()
			return
		}
		Claims = claims
		c.Next()
	}
}
