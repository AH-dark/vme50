package main

import (
	"flag"
	"github.com/AH-dark/random-donate/bootstrap"
	"github.com/AH-dark/random-donate/model"
	"github.com/AH-dark/random-donate/pkg/conf"
	"github.com/AH-dark/random-donate/pkg/utils"
	"github.com/AH-dark/random-donate/router"
)

func init() {
	flag.StringVar(&conf.FilePath, "c", utils.RelativePath("conf.ini"), "配置文件路径")
	flag.BoolVar(&conf.UpdateDatabase, "u", false, "是否更新数据库")
	flag.Parse()

	bootstrap.Init()
}

func main() {
	// 数据库初始化
	model.Init()

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
