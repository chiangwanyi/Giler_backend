package main

import (
	"giler-backend/config"
	"giler-backend/server"
	"log"
	"os"
)

func main(){
	config.Init()

	router := server.NewRouter()

	if err := router.Run(os.Getenv("HTTP_ADDR")); err != nil {
		log.Fatal("开启 HTTP 服务器失败：", err)
	}
}
