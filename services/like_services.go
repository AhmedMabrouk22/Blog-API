package services

import (
	"main/models"
	"main/repositories"
)

type LikeServices interface {
	Add(blogId uint, userId uint) error
	Delete(blogId uint, userId uint) error
	Get(blogId uint) ([]models.LikeRes, error)
}

type LikeServicesImp struct {
	LikeRepository repositories.LikeRepository
}

// Get implements LikeServices.
func (l *LikeServicesImp) Get(blogId uint) ([]models.LikeRes, error) {
	return l.LikeRepository.Get(blogId)
}

// Add implements LikeServices.
func (l *LikeServicesImp) Add(blogId uint, userId uint) error {
	return l.LikeRepository.Add(blogId, userId)
}

// Delete implements LikeServices.
func (l *LikeServicesImp) Delete(blogId uint, userId uint) error {
	return l.LikeRepository.Delete(blogId, userId)
}

func NewLikeServices(LikeRepository repositories.LikeRepository) LikeServices {
	return &LikeServicesImp{LikeRepository: LikeRepository}
}
