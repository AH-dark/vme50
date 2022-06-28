package database

import "gorm.io/gorm"

type DonateInfo struct {
	gorm.Model
	Name    string `gorm:"not null" json:"name"`
	Email   string `gorm:"unique;not null;index:donateInfo_email" json:"email"`
	Payment string `gorm:"not null" json:"payment"`
	Url     string `gorm:"not null" json:"url"`
}
