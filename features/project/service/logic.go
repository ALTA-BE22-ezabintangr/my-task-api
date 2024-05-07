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
	if input.ProjectName == "" || input.UserID == 0 {
		return errors.New("nama project/userID tidak boleh kosong")
	}
	err := p.projectData.Insert(input)
	if err != nil {
		return err
	}

	return nil
}

// GetAll implements project.ServiceInterface.
func (p *projectService) GetAll() ([]project.Core, error) {
	return p.projectData.SelectAll()
}

// Update implements project.ServiceInterface.
func (p *projectService) Update(id uint, input project.Core) error {
	return p.projectData.Update(id, input)
}

// Delete implements project.ServiceInterface.
func (p *projectService) Delete(id uint) error {
	panic("unimplemented")
}
