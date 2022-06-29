package service

import "github.com/AH-dark/random-donate/model"

func Count(info interface{}, args ...interface{}) (int64, error) {
	var count int64
	err := model.DB.Model(info).Where(args).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}
