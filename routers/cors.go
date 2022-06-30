package routers

import (
	"github.com/AH-dark/random-donate/pkg/conf"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func initCors(r *gin.Engine) {
	config := cors.Config{
		AllowOriginFunc: func(origin string) bool {
			if gin.Mode() != gin.ReleaseMode {
				return true
			}

			for _, v := range conf.CORSConfig.AllowOrigins {
				if v == origin {
					return true
				}
			}

			return false
		},
		AllowMethods:     conf.CORSConfig.AllowMethods,
		AllowHeaders:     conf.CORSConfig.AllowHeaders,
		AllowCredentials: true,
		ExposeHeaders:    conf.CORSConfig.ExposeHeaders,
		MaxAge:           3600,
	}

	handler := cors.New(config)

	r.Use(handler)
}
