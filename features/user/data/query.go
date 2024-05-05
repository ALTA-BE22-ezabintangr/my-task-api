package data

import (
	"myTaskApp/features/user"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.DataInterface {
	return &userQuery{
		db: db,
	}
}

// Insert implements user.DataInterface.
func (u *userQuery) Insert(input user.Core) error {
	var userGorm User

	userGorm = User{
		Model:    gorm.Model{},
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Phone:    input.Phone,
		Address:  input.Address,
	}
	tx := u.db.Create(&userGorm)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// SelectAll implements user.DataInterface.
func (u *userQuery) SelectAll() ([]user.Core, error) {
	panic("unimplemented")
	// var allUser []User
	// tx := u.db.Find(&allUser)
	// if tx.Error != nil {
	// 	return nil, tx.Error
	// }

}

// Delete implements user.DataInterface.
func (u *userQuery) Delete(id uint) error {
	panic("unimplemented")
}
