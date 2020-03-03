package api

import (
	"giler-backend/modules"
	"github.com/gin-gonic/gin"
)


func ConnWebSocket(c *gin.Context) {
	modules.ServeWs(c.Writer, c.Request)
}
