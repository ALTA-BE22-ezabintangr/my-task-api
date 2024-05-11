package data

import (
	"errors"
	"myTaskApp/features/project"
	"myTaskApp/features/task/data"

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
func (p *projectQuery) GetProjectById(id uint) (project.Core, error) {
	var projectId Project
	tx := p.db.First(&projectId, id)
	if tx.Error != nil {
		return project.Core{}, tx.Error
	}

	var getTask []data.Task
	tx2 := p.db.Joins("JOIN projects ON projects.id = tasks.project_id").Where("projects.id = ?", projectId.ID).Find(&getTask)
	if tx2.Error != nil {
		return project.Core{}, tx2.Error
	}

	var resultGetTask []project.TaskListResponseCore
	for _, v := range getTask {
		resultGetTask = append(resultGetTask, project.TaskListResponseCore{
			ID:              v.ID,
			TaskName:        v.TaskName,
			DescriptionTask: v.DescriptionTask,
			StatusTask:      v.StatusTask,
		})
	}

	projectIdCore := project.Core{
		ProjectName: projectId.ProjectName,
		Description: projectId.Description,
		TaskList:    resultGetTask,
	}

	return projectIdCore, nil

}

// Update implements project.DataInterface.
func (p *projectQuery) Update(id uint, input project.Core) error {
	var projectCurrent Project
	tx := p.db.First(&projectCurrent, id)
	if tx.Error != nil {
		return tx.Error
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

// GetUserByProjectId implements project.DataInterface.
func (p *projectQuery) GetUserByProjectId(id uint) (project.Core, error) {
	var projectUserId Project
	tx := p.db.First(&projectUserId, id)
	if tx.Error != nil {
		return project.Core{}, tx.Error
	}
	projectIdCore := project.Core{
		ID:     id,
		UserID: projectUserId.UserID,
	}

	return projectIdCore, nil
}
