package user

import "time"

type Core struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	Phone     string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DataInterface interface {
	Insert(input Core) error
	SelectAll() ([]Core, error)
	Delete(id uint) error
	Update(id uint, input Core) error
	Login(email, password string) (*Core, error)
}

type ServiceInterface interface {
	Create(input Core) error
	GetAll() ([]Core, error)
	Delete(id uint) error
	Update(id uint, input Core) error
	Login(email, password string) (data *Core, token string, err error)
}
