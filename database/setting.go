package database

import (
	"github.com/AH-dark/random-donate/pkg/cache"
	"gorm.io/gorm"
	"strconv"
)

type Setting struct {
	gorm.Model
	Name  string `gorm:"not null;unique;index:setting_key" json:"name"`
	Type  string `gorm:"not null" json:"type"`
	Value string `gorm:"size:65535" json:"value"`
}

func IsTrue(value string) bool {
	return value == "1" || value == "true"
}

// GetSettingByName 用 Name 获取设置值
func GetSettingByName(name string) string {
	return GetSettingByNameFromTx(DB, name)
}

// GetSettingByNameFromTx 用 Name 获取设置值，使用事务
func GetSettingByNameFromTx(tx *gorm.DB, name string) string {
	var setting Setting

	// 优先从缓存中查找
	cacheKey := "setting_" + name
	if optionValue, ok := cache.Get(cacheKey); ok {
		return optionValue.(string)
	}

	// 尝试数据库中查找
	if tx == nil {
		tx = DB
		if tx == nil {
			return ""
		}
	}

	result := tx.Where("name = ?", name).First(&setting)
	if result.Error == nil {
		_ = cache.Set(cacheKey, setting.Value, -1)
		return setting.Value
	}

	return ""
}

// GetSettingByNameWithDefault 用 Name 获取设置值, 取不到时使用缺省值
func GetSettingByNameWithDefault(name, fallback string) string {
	res := GetSettingByName(name)
	if res == "" {
		return fallback
	}
	return res
}

// GetSettingByNames 用多个 Name 获取设置值
func GetSettingByNames(names ...string) map[string]string {
	var queryRes []Setting
	res, miss := cache.GetSettings(names, "setting_")

	if len(miss) > 0 {
		DB.Where("name IN (?)", miss).Find(&queryRes)
		for _, setting := range queryRes {
			res[setting.Name] = setting.Value
		}
	}

	_ = cache.SetSettings(res, "setting_")
	return res
}

// GetSettingByType 获取一个或多个分组的所有设置值
func GetSettingByType(types []string) map[string]string {
	var queryRes []Setting
	res := make(map[string]string)

	DB.Where("type IN (?)", types).Find(&queryRes)
	for _, setting := range queryRes {
		res[setting.Name] = setting.Value
	}

	return res
}

// GetIntSetting 获取整形设置值，如果转换失败则返回默认值defaultVal
func GetIntSetting(key string, defaultVal int) int {
	res, err := strconv.Atoi(GetSettingByName(key))
	if err != nil {
		return defaultVal
	}
	return res
}
