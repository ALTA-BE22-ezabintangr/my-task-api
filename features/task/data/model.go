package data

import (
	dataproject "myTaskApp/features/project/data"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ProjectID       uint
	TaskName        string
	DescriptionTask string
	StatusTask      string
	Projects        dataproject.Project `gorm:"foreignkey:ProjectID"`
}
