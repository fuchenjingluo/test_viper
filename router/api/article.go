package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
	"test_viper/config"
	"test_viper/models"
	"test_viper/utils"
	"time"
)

func CreateArticle(c *gin.Context) {
	var article models.Article
	article.Title = getParam(c, "title")
	article.Content = getParam(c, "content")

	if err := utils.Validate.Struct(article); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.(validator.ValidationErrors).Translate(utils.Trans)})
		return
	}

	if err := models.CreatArticle(article); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	pattern := "article:*"                             // 匹配所有以 article:title: 开头的键
	keys, err := config.RedisDB.Keys(pattern).Result() // 获取所有匹配的键
	if err != nil {
		// 处理错误
		fmt.Println("Error retrieving keys:", err)
		return
	}
	fmt.Println(keys)
	for _, key := range keys {
		config.RedisDB.Del(key) // 删除匹配的键
	}
	if article.Title != "" {
		pattern := "article:title:*"                       // 匹配所有以 article:title: 开头的键
		keys, err := config.RedisDB.Keys(pattern).Result() // 获取所有匹配的键
		if err != nil {
			// 处理错误
			fmt.Println("Error retrieving keys:", err)
			return
		}
		fmt.Println(keys)
		for _, key := range keys {
			config.RedisDB.Del(key) // 删除匹配的键
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "创建成功",
		"data": article,
	})
}

func GetArticle(c *gin.Context) {
	var article models.Article
	article.Title = getParam(c, "title")
	id, _ := strconv.Atoi(getParam(c, "id"))
	article.ID = uint(id)

	var cachedData string

	// 根据 ID 获取缓存
	if article.ID > 0 {
		cacheKey := fmt.Sprintf("article:%d", article.ID)
		cachedData, _ = config.RedisDB.Get(cacheKey).Result()
	}

	// 如果缓存为空，根据标题获取缓存
	if cachedData == "" && article.Title != "" {
		cacheKey := fmt.Sprintf("article:title:%s", article.Title)
		cachedData, _ = config.RedisDB.Get(cacheKey).Result()
	}

	if cachedData != "" { // 如果缓存存在
		var articles []models.Article
		json.Unmarshal([]byte(cachedData), &articles) // 反序列化
		c.JSON(http.StatusOK, gin.H{
			"code":  200,
			"msg":   "查询成功",
			"data":  articles,
			"total": len(articles),
		})
		return
	}

	// 查询数据库
	articles, total, err := models.GetArticle(article)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 将查询结果存入 Redis 缓存
	if total > 0 {
		// 序列化为 JSON 字符串
		dataToCache, _ := json.Marshal(articles)

		// 根据 ID 缓存
		if article.ID > 0 {
			cacheKey := fmt.Sprintf("article:%d", article.ID)
			config.RedisDB.Set(cacheKey, dataToCache, 10*time.Second)
		}

		// 根据标题缓存
		if article.Title != "" {
			cacheKey := fmt.Sprintf("article:title:%s", article.Title)
			config.RedisDB.Set(cacheKey, dataToCache, 10*time.Minute)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"msg":   "查询成功",
		"data":  articles,
		"total": total,
	})
}

func LikeArticle(c *gin.Context) {
	var articleID string
	articleID = getParam(c, "articleID")
	if articleID == "" {
		c.JSON(http.StatusOK, gin.H{
			"msg": "文章ID不能为空",
		})
		return
	}
	if err := models.LikeArticle(articleID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "点赞成功",
		"data": articleID,
	})
}

func GetArticleLikes(c *gin.Context) {
	var articleID string
	articleID = getParam(c, "articleID")
	if articleID == "" {
		c.JSON(http.StatusOK, gin.H{
			"msg": "文章ID不能为空",
		})
		return
	}
	likes, err := models.GetArticleLikes(articleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": likes,
	})
}
