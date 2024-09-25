package router

import (
	"github.com/gin-gonic/gin"
	"test_viper/moddleware/jwt"
	"test_viper/router/api"
	"test_viper/utils"
)

func SetupRouter() *gin.Engine {
	utils.InitTrans()
	r := gin.Default()
	auth := r.Group("api/auth")
	{
		auth.POST("/login", api.Login)
		auth.POST("/registered", api.Registered)
	}
	r.GET("/test", jwt.Jwt(), api.Test)
	r.POST("/CreatArticle", jwt.Jwt(), api.CreateArticle)
	r.POST("/GetArticle", jwt.Jwt(), api.GetArticle)
	r.POST("/LikeArticle", jwt.Jwt(), api.LikeArticle)
	r.POST("/GetArticleLikes", jwt.Jwt(), api.GetArticleLikes)
	return r
}
