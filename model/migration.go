package model

import (
	"github.com/AH-dark/random-donate/pkg/conf"
	"github.com/AH-dark/random-donate/pkg/encrypt"
	"github.com/AH-dark/random-donate/pkg/utils"
	"gorm.io/gorm"
)

func needMigration() bool {
	var str Setting
	err := DB.Where(&Setting{
		Name: "db_version",
		Type: "system",
	}).First(&str).Error

	if err != nil {
		return true
	}

	if str.Value != conf.DbVersion {
		return true
	}

	return false
}

func migration() {
	if !conf.UpdateDatabase && !needMigration() {
		utils.Log().Info("数据库版本匹配，跳过初始化")
		return
	}

	utils.Log().Info("准备进行数据库初始化")

	// 自动迁移模式
	if conf.DatabaseConfig.Type == "mysql" {
		DB = DB.Set("gorm:table_options", "ENGINE=InnoDB")
	}

	err := DB.AutoMigrate(&Setting{}, &DonateInfo{}, &User{})
	if err != nil {
		utils.Log().Panic("数据库初始化时错误，%s", err.Error())
		return
	}

	addDefaultSettings()
	addDefaultUser()

	utils.Log().Info("准备更新数据库版本信息")

	updateSetting("app_version", conf.AppVersion)
	updateSetting("db_version", conf.DbVersion)

	utils.Log().Info("数据库初始化结束")
}

func addDefaultSettings() {
	for _, value := range defaultSettings {
		DB.Where(&Setting{
			Name: value.Name,
		}).Create(&value)
	}
}

func addDefaultUser() {
	var count int64
	DB.Model(&User{}).Where(&User{
		Model: gorm.Model{
			ID: 1,
		},
	}).Count(&count)

	if count != 0 {
		return
	}

	pass := utils.RandStringRunes(12)
	data := &User{
		Username:    "admin",
		Nickname:    "Admin",
		Email:       "admin@example.com",
		Password:    encrypt.Pass(pass),
		Description: "",
		Role:        1,
	}

	DB.Model(&User{}).Where(&User{
		Model: gorm.Model{
			ID: 1,
		},
	}).Create(&data)

	utils.Log().Info("初始管理员用户名：%v", data.Username)
	utils.Log().Info("初始管理员邮箱：%v", data.Email)
	utils.Log().Info("初始管理员密码：%v", pass)
}

func updateSetting(name string, value string) {
	var data Setting
	DB.Model(&Setting{}).Where(&Setting{
		Name: name,
	}).First(&data)

	data.Value = value

	DB.Save(&data)
}
