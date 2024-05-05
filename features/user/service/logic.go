package service

import (
	"errors"
	"myTaskApp/features/user"
)

type userService struct {
	userData user.DataInterface
}

func New(ud user.DataInterface) user.ServiceInterface {
	return &userService{
		userData: ud,
	}
}

// Create implements user.ServiceInterface.
func (u *userService) Create(input user.Core) error {
	if input.Name == "" || input.Email == "" || input.Password == "" {
		return errors.New("nama/email/password tidak boleh kosong")
	}

	err := u.userData.Insert(input)
	if err != nil {
		return err
	}
	return nil
}

// GetAll implements user.ServiceInterface.
func (u *userService) GetAll() ([]user.Core, error) {
	return u.userData.SelectAll()
}
