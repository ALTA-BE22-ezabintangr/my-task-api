package data

import (
	_userData "myTaskApp/features/user/data"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	UserID      uint
	ProjectName string
	Description string
	User        _userData.User `gorm:"references:ID;foreignKey:UserID"`
}
