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

// func GetProjectsForUser(db *gorm.DB, userID int) ([]Project, error) {
// 	var projects []Project
// 	err := db.Preload("User").Where("user_id = ?", userID).Find(&projects).Error
// 	return projects, err
// }
