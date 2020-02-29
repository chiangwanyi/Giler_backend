package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"regexp"
)

// Cors 跨域配置
func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie"}
	if gin.Mode() == gin.ReleaseMode {
		// 生产环境需要配置跨域域名，否则403
		config.AllowOrigins = []string{os.Getenv("DOMAIN_NAME")}
	} else {
		// 测试环境下模糊匹配本地开头的请求
		config.AllowOriginFunc = func(origin string) bool {
			if regexp.MustCompile(`^http://127\.0\.0\.1`).MatchString(origin) {
				return true
			}
			if regexp.MustCompile(`^http://localhost`).MatchString(origin) {
				return true
			}
			if regexp.MustCompile(`^http://192.168`).MatchString(origin) {
				return true
			}
			return false
		}
	}
	config.AllowCredentials = true
	return cors.New(config)
}
