package data

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ProjectID       uint
	TaskName        string
	DescriptionTask string
	StatusTask      string
	// Projects        data.Project `gorm:"foreignkey:ProjectID"`
}
