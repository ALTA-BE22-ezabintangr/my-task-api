package data

import (
	"myTaskApp/features/task/data"
	userData "myTaskApp/features/user/data"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	UserID      uint
	ProjectName string
	Description string
	User        userData.User `gorm:"foreignKey:UserID"`
	Tasks       []data.Task   `gorm:"foreignKey:ProjectID"`
}
