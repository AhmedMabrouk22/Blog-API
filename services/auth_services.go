package services

import (
	"errors"
	"main/models"
	"main/repositories"
	"main/utils"
)

type AuthServices interface {
	Login(email, password string) (models.User, error)
	SignUp(user models.User) (models.User, error)
	ChangePassword(id interface{}, curPassword, newPassword string) error
}

type authServicesImp struct {
	userRepository repositories.UserRepository
}

// ChangePassword implements AuthServices.
func (a *authServicesImp) ChangePassword(id interface{}, curPassword, newPassword string) error {

	// 1) If the curPassword is correct
	user, err := a.userRepository.FindUserById(id)
	if err != nil {
		return err
	}

	err = utils.VerifyPassword(user.Password, curPassword)
	if err != nil {
		return errors.New("Please enter the correct password")
	}

	// 2) Hash new Password
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}

	// 3) Change password
	err = a.userRepository.UpdatePassword(id, hashedPassword)
	if err != nil {
		return err
	}
	return nil
}

// SignUp implements AuthServices.
func (a *authServicesImp) SignUp(user models.User) (models.User, error) {
	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return models.User{}, err
	}

	newUser := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: hashedPassword,
		Image:    user.Image,
	}

	result, err := a.userRepository.CreateUser(newUser)
	if err != nil {
		return models.User{}, err
	}

	return result, nil

}

// Login implements AuthServices.
func (a *authServicesImp) Login(email string, password string) (models.User, error) {
	user, err := a.userRepository.FindUserByEmail(email)

	if err != nil {
		return models.User{}, errors.New("Invalid email or password")
	}

	err = utils.VerifyPassword(user.Password, password)
	if err != nil {
		return models.User{}, errors.New("Invalid email or password")
	}

	return user, nil
}

func NewAuthServices(userRepository repositories.UserRepository) AuthServices {
	return &authServicesImp{userRepository: userRepository}
}
