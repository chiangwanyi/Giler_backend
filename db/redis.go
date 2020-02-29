package db

import (
	"log"
	"os"

	"github.com/go-redis/redis/v7"
)

// RedisConn Redis 连接
var RedisConn *redis.Client

// InitRedisConnect 初始化 Redis 连接
func InitRedisConnect() {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PWD"),
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal("初始化 Redis 失败：", err)
	}
	RedisConn = client
}
