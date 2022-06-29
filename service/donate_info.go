package service

import (
	"errors"
	"github.com/AH-dark/random-donate/model"
	"math/rand"
)

func DonateInfoCount(info *model.DonateInfo) (int64, error) {
	var count int64
	err := model.DB.Where(info).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func DonateInfoIsExist(info *model.DonateInfo) (bool, error) {
	count, err := DonateInfoCount(info)
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

func DonateInfoRandomGet() (model.DonateInfo, error) {
	count, err := DonateInfoCount(&model.DonateInfo{})
	if err != nil {
		return model.DonateInfo{}, err
	}
	if count < 1 {
		return model.DonateInfo{}, errors.New("no data exist")
	}

	data := model.DonateInfo{}
	err = model.DB.Offset(rand.Intn(int(count - 1))).First(&data).Error
	if err != nil {
		return model.DonateInfo{}, err
	}

	return data, nil
}
