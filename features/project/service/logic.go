package service

import (
	"errors"
	"myTaskApp/features/project"
)

type projectService struct {
	projectData project.DataInterface
}

func New(pd project.DataInterface) project.ServiceInterface {
	return &projectService{
		projectData: pd,
	}
}

// Create implements project.ServiceInterface.
func (p *projectService) Create(input project.Core) error {
	if input.ProjectName == "" {
		return errors.New("nama project/userID tidak boleh kosong")
	}
	err := p.projectData.Insert(input)
	if err != nil {
		return err
	}

	return nil
}

// GetAll implements project.ServiceInterface.
func (p *projectService) GetAll(id uint) ([]project.Core, error) {
	return p.projectData.SelectAll(id)
}

// GetProjectById implements project.ServiceInterface.
func (p *projectService) GetProjectById(id uint, idUser uint) (input project.Core, err error) {
	result, err2 := p.projectData.GetUserByProjectId(id)
	if err2 != nil {
		return project.Core{}, err2
	}
	if result.UserID != idUser {
		return project.Core{}, errors.New("id project tidak sesuai dengan milik anda")
	}
	return p.projectData.GetProjectById(id)
}

// Update implements project.ServiceInterface.
func (p *projectService) Update(id uint, idUser uint, input project.Core) error {
	if input.UserID != idUser {
		return errors.New("id project bukan milik anda")
	}
	return p.projectData.Update(id, input)
}

// Delete implements project.ServiceInterface.
func (p *projectService) Delete(id uint, idUser uint) error {
	return p.projectData.Delete(id, idUser)
}
