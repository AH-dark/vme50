package service

import (
	"errors"
	"github.com/AH-dark/random-donate/model"
	"github.com/AH-dark/random-donate/pkg/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const sessNamespace = "user_info"

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

func GetUserBySession(c *gin.Context) (*model.User, error) {
	var user model.User

	sess := utils.GetSession(c, sessNamespace)
	if sess == nil {
		return nil, errors.New("session not exist")
	}

	err := model.DB.Model(&model.User{}).Where(&model.User{
		Model: gorm.Model{
			ID: sess.(uint),
		},
	}).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
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
