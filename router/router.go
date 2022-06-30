package router

import (
	"github.com/AH-dark/random-donate/pkg/conf"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// cors
	r.Use(cors.New(cors.Config{
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
	}))
	// gzip
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	// api
	initApiV1(r)

	return r
}
