package server

import (
	"giler-backend/api"
	"giler-backend/middlewares"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.Cors())

	user := r.Group("/api/v1")
	{
		user.POST("user", api.CreateUser)
		user.GET("user", api.QueryUser)
		user.GET("users", api.ListUser)
		user.PUT("user/:id", api.UpdateUser)
		user.DELETE("user/:id", api.DeleteUser)
	}

	return r
}
