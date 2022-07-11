package middleware

import "github.com/gin-gonic/gin"

// CacheControl 缓存控制中间件
func CacheControl() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "private, no-cache")
		c.Next()
	}
}
