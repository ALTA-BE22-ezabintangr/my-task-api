package data

import (
	"myTaskApp/features/project"

	"gorm.io/gorm"
)

type projectQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) project.DataInterface {
	return &projectQuery{
		db: db,
	}
}

// Insert implements project.DataInterface.
func (p *projectQuery) Insert(input project.Core) error {
	var projectGorm Project

	projectGorm = Project{
		Model:       gorm.Model{},
		UserID:      input.UserID,
		ProjectName: input.ProjectName,
		Description: input.Description,
	}

	tx := p.db.Create(&projectGorm)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// SelectAll implements project.DataInterface.
func (p *projectQuery) SelectAll() ([]project.Core, error) {
	var allProject []Project
	tx := p.db.Find(&allProject)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var allProjectCore []project.Core
	for _, v := range allProject {
		allProjectCore = append(allProjectCore, project.Core{
			ID:          v.ID,
			UserID:      v.UserID,
			ProjectName: v.ProjectName,
			Description: v.Description,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		})
	}

	return allProjectCore, nil
}

// Update implements project.DataInterface.
func (p *projectQuery) Update(id uint, input project.Core) error {
	tx := p.db.Model(&Project{}).Where("id = ?", id).Updates(input)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Delete implements project.DataInterface.
func (p *projectQuery) Delete(id uint) error {
	tx := p.db.Delete(&Project{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
