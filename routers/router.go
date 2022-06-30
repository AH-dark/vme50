package routers

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// cors
	initCors(r)

	// gzip
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	// api
	initApiV1(r)

	return r
}
