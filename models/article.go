package models

import (
	"gorm.io/gorm"
	"test_viper/config"
)

type Article struct {
	gorm.Model
	Title   string `json:"title" validate:"required,max=100"`
	Content string `json:"content" validate:"required"`
	Preview string `json:"preview"`
	Likes   int    `json:"likes"`
}

func CreatArticle(article Article) error {
	if err := config.DB.AutoMigrate(&Article{}); err != nil {
		return err
	}
	if err := config.DB.Create(&article).Error; err != nil {
		return err
	}
	return nil
}
func GetArticle(article Article) ([]Article, int64, error) {
	query := config.DB.Model(&Article{})
	var articles []Article // 声明一个切片来存储查询结果
	// 如果有 ID，则根据 ID 查询
	if article.ID > 0 {
		query = query.Where("id = ?", article.ID)
	}
	// 如果有 Title，则根据 Title 查询
	if article.Title != "" {
		query = query.Where("title LIKE ?", "%"+article.Title+"%")
	}
	// 执行查询
	if err := query.Find(&articles).Error; err != nil {
		return nil, 0, err
	}
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return articles, total, nil
}
func LikeArticle(articleID string) error {
	if articleID == "" {
		return nil
	}
	likeKey := "article:" + articleID + ":likes"
	if err := config.RedisDB.Incr(likeKey).Err(); err != nil {
		return err
	}
	return nil
}
func GetArticleLikes(articleID string) (string, error) {
	likeKey := "article:" + articleID + ":likes"
	likes, err := config.RedisDB.Get(likeKey).Result()
	if err != nil {
		return "", err
	}
	return likes, nil
}
