package router

import (
	"github.com/AH-dark/random-donate/controller"
	"github.com/AH-dark/random-donate/middleware"
	"github.com/gin-gonic/gin"
)

func initApiV1(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	v1.Use(middleware.Session())

	// ping
	v1.GET("/ping", controller.PingHandler)

	// settings
	{
		v1.GET("/siteInfo", controller.GetBasicSettings)
		v1.GET("/settings/basic", controller.GetBasicSettings)
	}

	// donate
	{
		v1.GET("/donate/random", controller.DonateRandomGetHandler)
		v1.POST("/donate", controller.DonatePostHandler)
	}
}
