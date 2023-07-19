package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type Package struct {
	gorm.Model
	PackageName string `gorm:"size:255;not null" json:"packageName"`
	UserId      int    `gorm:"not null;" json:"userId"`
}

func (_package *Package) CreatePackage() (*Package, error) {
	var err error
	err = DB.Create(&_package).Error
	if err != nil {
		return &Package{}, err
	}
	return _package, nil
}

func GetTotalPackage(userId uint) (int, error) {
	var count int
	if err := DB.Model(&Package{}).Where("user_id = ?", userId).Count(&count).Error; err != nil {
		return 0, errors.New("No Package associate with this user")
	}

	return count, nil
}

func GetAllPackage(userId uint) ([]Package, error) {
	var _package []Package
	if err := DB.Where("user_id = ?", userId).Find(&_package).Error; err != nil {
		return []Package{}, errors.New("No Package associate with this user")
	}

	return _package, nil
}
