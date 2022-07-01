package bootstrap

import (
	"fmt"
	"github.com/AH-dark/random-donate/pkg/conf"
	"github.com/gin-gonic/gin"
	"io/fs"
)

func Init(static fs.FS) {
	fmt.Println("应用程序已启动")

	// init config
	conf.Init()

	// init static
	InitStatic(static)

	// Debug 关闭时，切换为生产模式
	if !conf.SystemConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
}
