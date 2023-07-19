package models

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"
)

type Image struct {
	gorm.Model
	ImageUrl string `gorm:"not null;unique" json:"imageUrl"`
	UserId   int    `gorm:"not null;" json:"userId"`
}

type RelatedImage struct {
	Id         int
	RelativeId int
	ImageUrl   string
}

func (image *Image) SaveImage() (*Image, error) {
	var err error
	err = DB.Create(&image).Error
	if err != nil {
		return &Image{}, err
	}
	return image, nil
}

func GetImage(userId uint) ([]Image, error) {
	var res []Image = []Image{}

	var image []Image
	if err := DB.Where("user_id = ?", userId).Find(&image).Error; err != nil {
		return []Image{}, errors.New("User not found!")
	}
	res = append(res, image...)

	log.Print(userId)
	rows, err := DB.Table("images").Select("images.id, family_members.relative_id, images.image_url").Joins("left join family_members on images.user_id = family_members.relative_id").Where("family_members.user_id = ?", userId).Rows()
	if err != nil {
		return res, nil
	}
	for rows.Next() {
		var relatedImage RelatedImage
		DB.ScanRows(rows, &relatedImage)

		var i Image
		i.ID = uint(relatedImage.Id)
		i.UserId = relatedImage.RelativeId
		i.ImageUrl = relatedImage.ImageUrl
		res = append(res, i)
	}
	log.Print(res)

	return res, nil
}
