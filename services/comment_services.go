package services

import (
	"main/models"
	"main/repositories"
)

type CommentServices interface {
	Create(comment models.Comment) (models.CommentRes, error)
	Update(id interface{}, userId interface{}, comment models.Comment) error
	Delete(id interface{}, userId interface{}) error
	Get(blogId interface{}) ([]models.CommentRes, error)
}

type commentServicesImp struct {
	commentRepository repositories.CommentRepository
}

// Get implements CommentServices.
func (c *commentServicesImp) Get(blogId interface{}) ([]models.CommentRes, error) {
	return c.commentRepository.Get(blogId)
}

// Create implements CommentServices.
func (c *commentServicesImp) Create(comment models.Comment) (models.CommentRes, error) {
	return c.commentRepository.Create(comment)
}

// Delete implements CommentServices.
func (c *commentServicesImp) Delete(id interface{}, userId interface{}) error {
	return c.commentRepository.Delete(id, userId)
}

// Update implements CommentServices.
func (c *commentServicesImp) Update(id interface{}, userId interface{}, comment models.Comment) error {
	return c.commentRepository.Update(id, userId, comment)
}

func NewCommentServices(commentRepository repositories.CommentRepository) CommentServices {
	return &commentServicesImp{commentRepository: commentRepository}
}
