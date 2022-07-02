package service

import (
	"github.com/AH-dark/random-donate/model"
	"gorm.io/gorm"
)

func Login(login string, pass string) (model.User, error) {
	var user model.User

	err := model.DB.Model(&model.User{}).Where(&model.User{
		Email:    login,
		Password: pass,
	}).Or(&model.User{
		Username: login,
		Password: pass,
	}).First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func GetUserById(id uint) (model.User, error) {
	var user model.User

	err := model.DB.Model(&model.User{}).Where(&model.User{
		Model: gorm.Model{
			ID: id,
		},
	}).First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
