package router

import (
	"github.com/AH-dark/random-donate/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// cors
	r.Use(cors.Default())
	// gzip
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	// api
	v1 := r.Group("/api/v1")

	{
		v1.GET("/ping", controller.PingHandler)
	}

	return r
}
