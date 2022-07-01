package routers

import (
	"github.com/AH-dark/random-donate/controller"
	"github.com/AH-dark/random-donate/middleware"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// cors
	r.Use(middleware.Cors())

	// gzip
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	// statics
	r.Use(controller.FrontendHandler())

	// api
	initApiV1(r)

	return r
}
