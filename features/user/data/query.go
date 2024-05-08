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
	var allUsers []User
	tx := u.db.Find(&allUsers)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var allUserCore []user.Core
	for _, v := range allUsers {
		allUserCore = append(allUserCore, user.Core{
			ID:        v.ID,
			Name:      v.Name,
			Email:     v.Email,
			Password:  v.Password,
			Phone:     v.Phone,
			Address:   v.Address,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return allUserCore, nil
}

// SelectAllById implements user.DataInterface.
func (u *userQuery) SelectAllById() ([]user.Core, error) {
	panic("unimplemented")
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
func (u *userQuery) Login(email string, password string) (*user.Core, error) {
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

// var currentUser User
// tx := u.db.First(&currentUser, id)
// if tx.Error != nil {
// 	if tx.Error == gorm.ErrRecordNotFound {
// 		return errors.New("User not found")
// 	}
// 	return tx.Error
// }

// currentUser.Name = input.Name
// currentUser.Email = input.Email
// currentUser.Password = input.Password
// currentUser.Address = input.Address
// currentUser.Phone = input.Phone

// tx = u.db.Save(&currentUser)
// if tx.Error != nil {
// 	return tx.Error
// }
