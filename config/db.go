package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func InitDB() {

	dsn := AppConfig.Database.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(AppConfig.Database.MaxIdleConns) //设置数据库连接池中的最大空闲连接数
	sqlDB.SetMaxOpenConns(AppConfig.Database.MaxOpenConns) //设置数据库连接池中最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour)                    //每个连接的最大存活时间
	if err != nil {
		panic(fmt.Errorf("failed to connect database,err:%v\n", err))
	}
	DB = db
}
