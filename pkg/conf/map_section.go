package conf

import "github.com/go-playground/validator/v10"

// mapSection 将配置文件的 Section 映射到结构体上
func mapSection(section string, confStruct interface{}) error {
	err := conf.Section(section).MapTo(confStruct)
	if err != nil {
		return err
	}

	// 验证合法性
	validate := validator.New()
	err = validate.Struct(confStruct)
	if err != nil {
		return err
	}

	return nil
}
