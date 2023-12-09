package repositories

import (
	"errors"
	"main/models"
	"main/utils"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user models.User) (models.User, error)
	FindUserByEmail(email string) (models.User, error)
	FindUserById(id interface{}) (models.User, error)
	Update(id interface{}, user models.User) (models.User, error)
	UpdatePassword(id interface{}, password string) error
	GetUser(id interface{}) (models.UserResponse, error)
}

type userRepositoryImpl struct {
	Db *gorm.DB
}

// UpdatePassword implements UserRepository.
func (u *userRepositoryImpl) UpdatePassword(id interface{}, password string) error {
	user, err := u.FindUserById(id)
	if err != nil {
		return err
	}

	user.Password = password
	result := u.Db.Save(&user)
	return result.Error
}

// Update implements UserRepository.
func (u *userRepositoryImpl) Update(id interface{}, user models.User) (models.User, error) {

	curUser, err := u.FindUserById(id)
	if err != nil {
		return models.User{}, err
	}

	if user.Name != "" {
		curUser.Name = user.Name
	}

	if user.Image != "" {
		curUser.Image = user.Image
	}

	result := u.Db.Save(&curUser)
	if result.Error != nil {
		panic(result.Error.Error())
	}
	return curUser, nil
}

// FindUserById implements UserRepository.
func (u *userRepositoryImpl) FindUserById(id interface{}) (models.User, error) {
	var user models.User
	result := u.Db.First(&user, "id = ?", id)
	if result.Error != nil {
		return user, errors.New("invalid id")
	}

	return user, nil
}

// CreateUser implements UserRepository.
func (u *userRepositoryImpl) CreateUser(user models.User) (models.User, error) {
	result := u.Db.Create(&user)
	if result.Error != nil {
		return models.User{}, errors.New(result.Error.Error())
	}

	return user, nil
}

// FindUser implements UserRepository.
func (u *userRepositoryImpl) FindUserByEmail(email string) (models.User, error) {
	var user models.User
	result := u.Db.First(&user, "email = ?", email)

	if result.Error != nil {
		return user, errors.New("invalid email or password")
	}

	return user, nil

}

// GetUser implements UserRepository.
func (u *userRepositoryImpl) GetUser(id interface{}) (models.UserResponse, error) {
	var user models.User
	result := u.Db.Preload("Blogs.Topics").Preload("Blogs.Author").First(&user, id)
	if result.Error != nil {
		return models.UserResponse{}, result.Error
	}

	var userRes models.UserResponse
	userRes.Name = user.Name
	userRes.Image = user.Image

	var blogs []models.BlogResponse
	for _, val := range user.Blogs {
		blogs = append(blogs, utils.GetBlogRes(val))
	}

	userRes.Blogs = blogs

	return userRes, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{Db: db}
}
