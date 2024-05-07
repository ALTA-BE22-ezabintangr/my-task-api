package data

import (
	"myTaskApp/features/user/data"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	UserID      uint
	ProjectName string
	Description string
	User        data.User `gorm:"foreignKey:UserID"`
}
