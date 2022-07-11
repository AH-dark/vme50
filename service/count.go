package service

import (
	"errors"
	"github.com/AH-dark/random-donate/model"
	"gorm.io/gorm"
)

func Count(info interface{}, condition interface{}, args ...interface{}) (int64, error) {
	var count int64 = -1
	err := model.DB.Model(info).Where(condition, args...).Count(&count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return count, err
	}

	return count, nil
}
