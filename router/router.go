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

		{
			v1.GET("/siteInfo", controller.GetBasicSettings)
			v1.GET("/settings/basic", controller.GetBasicSettings)
		}

		{
			v1.GET("/donate/random", controller.DonateRandomGetHandler)
			v1.POST("/donate", controller.DonatePostHandler)
		}
	}

	return r
}
