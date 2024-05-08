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
	SelectAll() ([]Core, error)
	GetProjectById(id uint) (Core, error)
	Delete(id uint) error
	Update(id uint, input Core) error
}

type ServiceInterface interface {
	Create(input Core) error
	GetAll() ([]Core, error)
	GetProjectById(id uint) (Core, error)
	Delete(id uint) error
	Update(id uint, input Core) error
}
