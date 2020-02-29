package server

import (
	"giler-backend/middlewares"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.Cors())

	return r
}
