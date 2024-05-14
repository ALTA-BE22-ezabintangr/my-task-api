package service

import (
	"errors"
	"myTaskApp/features/project"
	"myTaskApp/features/user"
)

type projectService struct {
	projectData project.DataInterface
	userData    user.DataInterface
}

func New(pd project.DataInterface, ud user.DataInterface) project.ServiceInterface {
	return &projectService{
		projectData: pd,
		userData:    ud,
	}
}

// Create implements project.ServiceInterface.
func (p *projectService) Create(input project.Core) error {
	result, err := p.userData.SelectProfileById(input.UserID)
	if err != nil {
		return err
	}

	if result.ID != input.UserID {
		return errors.New("user not found")
	}
	if input.ProjectName == "" {
		return errors.New("nama project tidak boleh kosong")
	}

	err2 := p.projectData.Insert(input)
	if err2 != nil {
		return err2
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
	result, err2 := p.projectData.GetUserByProjectId(id)
	if err2 != nil {
		return err2
	}
	if result.UserID != idUser {
		return errors.New("id project bukan milik anda")
	}
	return p.projectData.Update(id, input)
}

// Delete implements project.ServiceInterface.
func (p *projectService) Delete(id uint, idUser uint) error {
	result, err := p.projectData.GetUserByProjectId(id)
	if err != nil {
		return err
	}
	if result.UserID != idUser {
		return errors.New("id project bukan milik anda")
	}
	return p.projectData.Delete(id)
}
