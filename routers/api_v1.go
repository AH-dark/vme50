package routers

import (
	"github.com/AH-dark/random-donate/controller"
	"github.com/AH-dark/random-donate/middleware"
	"github.com/gin-gonic/gin"
)

func initApiV1(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	v1.Use(middleware.Session())
	v1.Use(middleware.CacheControl())

	// ping
	v1.GET("/ping", controller.PingHandler)

	// settings
	settings := v1.Group("settings")
	{
		v1.GET("siteInfo", controller.GetBasicSettings)
		settings.GET("basic", controller.GetBasicSettings)
	}

	// donate
	donate := v1.Group("donate")
	{
		donate.GET("random", controller.DonateRandomGetHandler)
		donate.POST("", controller.DonatePostHandler)
	}

	// user
	user := v1.Group("user")
	{
		user.GET("session", controller.SessionUserHandler)
		user.POST("login", controller.UserLogin)
	}
}
