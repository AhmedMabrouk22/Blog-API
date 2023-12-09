package services

import (
	"errors"
	"main/models"
	"main/repositories"
)

type UserServices interface {
	FindUserById(id interface{}) (models.User, error)
	Update(id interface{}, user models.User) (models.User, error)
	GetUser(id interface{}) (models.UserResponse, error)
}

type userServicesImp struct {
	userRepository repositories.UserRepository
}

// Update implements UserServices.
func (u *userServicesImp) Update(id interface{}, user models.User) (models.User, error) {
	updateUser, err := u.userRepository.Update(id, user)
	if err != nil {
		return models.User{}, errors.New("User not found")
	}

	return updateUser, nil
}

// FindUserById implements UserServices.
func (u *userServicesImp) FindUserById(id interface{}) (models.User, error) {
	user, err := u.userRepository.FindUserById(id)
	return user, err
}

// GetUser implements UserServices.
func (u *userServicesImp) GetUser(id interface{}) (models.UserResponse, error) {
	return u.userRepository.GetUser(id)
}

func NewUserServices(userRepository repositories.UserRepository) UserServices {
	return &userServicesImp{userRepository: userRepository}
}
