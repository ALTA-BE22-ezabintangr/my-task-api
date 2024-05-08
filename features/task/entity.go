package task

import (
	"time"
)

type Core struct {
	ID              uint
	UserID          uint
	ProjectID       uint
	TaskName        string
	DescriptionTask string
	StatusTask      string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type DataInterface interface {
	Insert(input Core) error
	Delete(id uint) error
	Update(id uint, input Core) error
	GetTaskById(id uint) (Core, error)
}

type ServiceInterface interface {
	Create(input Core) error
	Delete(id uint) error
	Update(id uint, input Core) error
	GetTaskbyId(id uint) (Core, error)
}
