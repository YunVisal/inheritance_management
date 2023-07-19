package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type FamilyMember struct {
	gorm.Model
	UserId     int  `gorm:"not null" json:"userId"`
	RelativeId int  `gorm:"not null;" json:"relativeId"`
	Relative   User `gorm:"foreignKey:RelativeID"`
}

type FamilyMemberResponse struct {
	Id         int    `json:"id"`
	RelativeId string `json:"relative_id"`
	Username   string `json:"username"`
}

func (family *FamilyMember) SaveRelationship() (*FamilyMember, error) {
	var err error
	err = DB.Create(&family).Error
	if err != nil {
		return &FamilyMember{}, err
	}

	var alternativeRelationship FamilyMember
	alternativeRelationship.UserId = family.RelativeId
	alternativeRelationship.RelativeId = family.UserId
	err = DB.Create(&alternativeRelationship).Error
	if err != nil {
		return &FamilyMember{}, err
	}

	return family, nil
}

func GetFamilyMember(userId uint) ([]FamilyMemberResponse, error) {
	var result []FamilyMemberResponse
	rows, err := DB.Table("users").Select("family_members.id, family_members.relative_id, users.username").Joins("left join family_members on users.id = family_members.relative_id").Where("family_members.user_id = ?", userId).Rows()
	if err != nil {
		return []FamilyMemberResponse{}, errors.New("No family member found")
	}
	for rows.Next() {
		var familyMember FamilyMemberResponse
		DB.ScanRows(rows, &familyMember)
		result = append(result, familyMember)
	}
	return result, nil
}
