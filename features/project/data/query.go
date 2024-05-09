package data

import (
	"errors"
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
func (p *projectQuery) SelectAll(id uint) ([]project.Core, error) {
	var allProject []Project
	tx := p.db.Where("user_id = ?", id).Find(&allProject)
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

// GetProjectById implements project.DataInterface.
func (p *projectQuery) GetProjectById(id uint, idUser uint) (project.Core, error) {
	var projectId Project
	tx := p.db.First(&projectId, id)
	if tx.Error != nil {
		return project.Core{}, tx.Error
	}

	if projectId.UserID != idUser {
		return project.Core{}, errors.New("id project tidak sesuai")
	}

	projectIdCore := project.Core{
		ID:          id,
		UserID:      projectId.UserID,
		ProjectName: projectId.ProjectName,
		Description: projectId.Description,
	}

	return projectIdCore, nil

}

// Update implements project.DataInterface.
func (p *projectQuery) Update(id uint, idUser uint, input project.Core) error {
	var projectCurrent Project
	tx := p.db.First(&projectCurrent, id)
	if tx.Error != nil {
		return tx.Error
	}

	if projectCurrent.UserID != idUser {
		return errors.New("id project bukan milik anda")
	}

	tx2 := p.db.Model(&Project{}).Where("id = ?", id).Updates(input)
	if tx2.Error != nil {
		return tx2.Error
	}

	return nil
}

// Delete implements project.DataInterface.
func (p *projectQuery) Delete(id uint, idUser uint) error {
	var projectDelete Project
	tx := p.db.First(&projectDelete, id)
	if tx.Error != nil {
		return tx.Error
	}

	if projectDelete.UserID != idUser {
		return errors.New("id project bukan milik anda")
	}

	tx2 := p.db.Delete(&projectDelete)
	if tx2.Error != nil {
		return tx2.Error
	}
	return nil
}
