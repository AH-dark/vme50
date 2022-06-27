package main

import (
	"flag"
	"github.com/AH-dark/random-donate/bootstrap"
	"github.com/AH-dark/random-donate/database"
	"github.com/AH-dark/random-donate/pkg/conf"
	"github.com/AH-dark/random-donate/pkg/utils"
	"github.com/AH-dark/random-donate/router"
)

var (
	confPath string
)

func init() {
	flag.StringVar(&confPath, "c", utils.RelativePath("conf.ini"), "配置文件路径")
	flag.Parse()

	bootstrap.Init(confPath)
}

func main() {
	// 数据库初始化
	database.Init()

	// 路由初始化
	r := router.InitRouter()

	// 监听
	utils.Log().Info("Application will listen " + conf.SystemConfig.Port + ".")
	err := r.Run(conf.SystemConfig.Port)
	if err != nil {
		utils.Log().Panic("Error when listen port "+conf.SystemConfig.Port+",", err.Error())
		return
	}
}
