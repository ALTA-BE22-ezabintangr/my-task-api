package data

import (
	projectData "myTaskApp/features/project/data"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ProjectID       uint
	TaskName        string
	DescriptionTask string
	StatusTask      string
	Projects        projectData.Project `gorm:"foreignkey:ProjectID"`
}
