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
	return p.projectData.GetProjectById(id, idUser)
}

// Update implements project.ServiceInterface.
func (p *projectService) Update(id uint, idUser uint, input project.Core) error {
	return p.projectData.Update(id, idUser, input)
}

// Delete implements project.ServiceInterface.
func (p *projectService) Delete(id uint, idUser uint) error {
	return p.projectData.Delete(id, idUser)
}
