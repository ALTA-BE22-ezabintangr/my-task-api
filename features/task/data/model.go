package data

import (
	projectData "myTaskApp/features/project/data"
	userData "myTaskApp/features/user/data"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	UserID          uint
	ProjectID       uint
	TaskName        string
	DescriptionTask string
	StatusTask      string
	User            userData.User       `gorm:"foreignKey:UserID"`
	Projects        projectData.Project `gorm:"foreignKey:ProjectID"`
}
