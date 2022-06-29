package conf

import (
	"github.com/AH-dark/random-donate/pkg/utils"
	"github.com/go-ini/ini"
)

var (
	FilePath       string
	UpdateDatabase bool
	conf           *ini.File
)

const defaultConf = `[System]
Debug = false
Listen = :8080
SessionSecret = {SessionSecret}
`

func Init() {
	var err error
	path := FilePath

	if path == "" || !utils.Exists(path) {
		// 创建初始配置文件
		confContent := utils.Replace(map[string]string{
			"{SessionSecret}": utils.RandStringRunes(64), // random
		}, defaultConf)

		f, err := utils.CreatNestedFile(path)
		if err != nil {
			utils.Log().Panic("无法创建配置文件, %s", err)
		}

		// 写入配置文件
		_, err = f.WriteString(confContent)
		if err != nil {
			utils.Log().Panic("无法写入配置文件, %s", err)
		}

		f.Close()
	}

	conf, err = ini.Load(path)
	if err != nil {
		utils.Log().Panic("无法解析配置文件 '%s': %s", path, err)
	}

	sections := map[string]interface{}{
		"Database": DatabaseConfig,
		"System":   SystemConfig,
		"CORS":     CORSConfig,
	}
	for sectionName, sectionStruct := range sections {
		err = mapSection(sectionName, sectionStruct)
		if err != nil {
			utils.Log().Panic("配置文件 %s 分区解析失败: %s", sectionName, err)
		}
	}

	// 映射数据库配置覆盖
	for _, key := range conf.Section("OptionOverwrite").Keys() {
		OptionOverwrite[key.Name()] = key.Value()
	}

	// 重设log等级
	if !SystemConfig.Debug {
		utils.Level = utils.LevelInformational
		utils.GlobalLogger = nil
		utils.Log()
	}
}
