package usersService

import (
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user User) (User, error)

	GetAllUsers() ([]User, error)

	UpdateUserByID(id uint, user User) (User, error)

	DeleteUserByID(id uint) error

	GetTasksForUser(uint) ([]User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user User) (User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) UpdateUserByID(id uint, user User) (User, error) {
	var existingUser User
	if err := r.db.First(&existingUser, id).Error; err != nil {
		return User{}, err
	}

	existingUser.Email = user.Email
	existingUser.Password = user.Password

	if err := r.db.Save(&existingUser).Error; err != nil {
		return User{}, err
	}

	return existingUser, nil

}

func (r *userRepository) DeleteUserByID(id uint) error {

	if err := r.db.Delete(&User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetTasksForUser(uint) ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}
