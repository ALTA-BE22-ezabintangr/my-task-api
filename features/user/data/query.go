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
func (u *userQuery) SelectProfileById(id uint) (*user.Core, error) {
	var userProfile User
	tx := u.db.First(&userProfile, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	userCore := user.Core{
		ID:        id,
		Name:      userProfile.Name,
		Email:     userProfile.Email,
		Password:  userProfile.Password,
		Phone:     userProfile.Phone,
		Address:   userProfile.Address,
		CreatedAt: userProfile.CreatedAt,
		UpdatedAt: userProfile.UpdatedAt,
	}

	return &userCore, nil
}

// Delete implements user.DataInterface.
func (u *userQuery) Delete(id uint) error {
	tx := u.db.Delete(&User{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Update implements user.DataInterface.
func (u *userQuery) Update(id uint, input user.Core) error {
	tx := u.db.Model(&User{}).Where("id=?", id).Updates(input)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// Login implements user.DataInterface.
func (u *userQuery) Login(email string) (*user.Core, error) {
	var userData User
	tx := u.db.Where("email = ?", email).First(&userData)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var userCore = user.Core{
		ID:        userData.ID,
		Name:      userData.Name,
		Email:     userData.Email,
		Password:  userData.Password,
		Phone:     userData.Phone,
		Address:   userData.Address,
		CreatedAt: userData.CreatedAt,
		UpdatedAt: userData.UpdatedAt,
	}

	return &userCore, nil
}
