package config

import "github.com/go-redis/redis"

var RedisDB *redis.Client

func InitRedis() {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "1438490700.a",
		DB:       0,
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		panic(err)
	}
	RedisDB = RedisClient
}
