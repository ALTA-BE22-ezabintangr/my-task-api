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
	SelectProfileById(id uint) (*Core, error)
	Delete(id uint) error
	Update(id uint, input Core) error
	Login(email string) (*Core, error)
}

type ServiceInterface interface {
	Create(input Core) error
	GetProfileUser(id uint) (*Core, error)
	Delete(id uint) error
	Update(id uint, input Core) error
	Login(email, password string) (data *Core, token string, err error)
}
