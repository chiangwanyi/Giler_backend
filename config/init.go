package config

import (
	"giler-backend/db"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	_ = godotenv.Load()

	gin.SetMode(os.Getenv("GIN_MODE"))

	db.InitMongoDBConnect()
	db.InitRedisConnect()

	//util.SetAccessToken()
}
