package db

import (
	"github.com/globalsign/mgo"
	"log"
	"os"
)

// MongoDBConn MongoDB 数据库连接实例
var MongoDBConn *mgo.Session

// InitMongoConnect 初始化 MongoDB 数据库
func InitMongoConnect() {
	m, err := mgo.Dial(os.Getenv("MONGODB_ADDR"))
	if err != nil {
		log.Fatal("初始化 MongoDB 数据库失败：", err)
	}
	MongoDBConn = m
}
