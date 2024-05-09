package project

import "time"

type Core struct {
	ID          uint
	UserID      uint
	ProjectName string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type DataInterface interface {
	Insert(input Core) error
	SelectAll(id uint) ([]Core, error)
	GetProjectById(id uint, idUser uint) (Core, error)
	Delete(id uint, idUser uint) error
	Update(id uint, idUser uint, input Core) error
}

type ServiceInterface interface {
	Create(input Core) error
	GetAll(id uint) ([]Core, error)
	GetProjectById(id uint, idUser uint) (Core, error)
	Delete(id uint, idUser uint) error
	Update(id uint, idUser uint, input Core) error
}
