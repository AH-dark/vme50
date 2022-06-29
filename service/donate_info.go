package service

import (
	"errors"
	"github.com/AH-dark/random-donate/model"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
)

func DonateInfoIsExist(info *model.DonateInfo) (bool, error) {
	count, err := Count(&info)
	if err != nil {
		return false, err
	}

	return count != 0, nil
}

func DonateInfoSave(info *model.DonateInfo) error {
	err := model.DB.Save(&info).Error
	return err
}

func DonateInfoFind(info *model.DonateInfo) (model.DonateInfo, error) {
	var dbData model.DonateInfo
	err := model.DB.Where(&info).First(&dbData).Error
	return dbData, err
}

func DonateInfoRandomGet(prevId string) (model.DonateInfo, error) {
	var data model.DonateInfo
	prevIdNum, _ := strconv.Atoi(prevId)
	not := model.DonateInfo{
		Model: gorm.Model{
			ID: uint(prevIdNum),
		},
	}

	var count int64

	err := model.DB.Model(&model.DonateInfo{}).Not(&not).Count(&count).Error
	if err != nil {
		return data, err
	}

	if count < 1 {
		return data, errors.New("no data exist")
	}

	err = model.DB.Not(&not).Offset(rand.Intn(int(count))).First(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}
