package repositories

import (
	"errors"
	"main/models"

	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(comment models.Comment) (models.CommentRes, error)
	Update(id interface{}, userId interface{}, comment models.Comment) error
	Delete(id interface{}, userId interface{}) error
	Get(blogId interface{}) ([]models.CommentRes, error)
}

type commentRepositoryImp struct {
	Db *gorm.DB
}

// Get implements CommentRepository.
func (c *commentRepositoryImp) Get(blogId interface{}) ([]models.CommentRes, error) {
	var blog models.Blog
	result := c.Db.Preload("Comments.User").First(&blog, blogId)
	if result.Error != nil {
		return []models.CommentRes{}, result.Error
	}

	var comments []models.CommentRes
	for _, val := range blog.Comments {
		comment := models.CommentRes{
			ID:      val.ID,
			Content: val.Content,
			Author: models.CommentAuthor{
				ID:    val.User.ID,
				Name:  val.User.Name,
				Image: val.User.Image,
			},
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

// Create implements CommentRepository.
func (c *commentRepositoryImp) Create(comment models.Comment) (models.CommentRes, error) {
	result := c.Db.Create(&comment)
	if result.Error != nil {
		return models.CommentRes{}, result.Error
	}

	var newComment models.CommentRes
	newComment.Content = comment.Content
	newComment.ID = comment.ID
	newComment.Author.ID = comment.User.ID
	newComment.Author.Name = comment.User.Name
	newComment.Author.Image = comment.User.Image

	return newComment, nil
}

// Delete implements CommentRepository.
func (c *commentRepositoryImp) Delete(id interface{}, userId interface{}) error {
	var curComment models.Comment
	result := c.Db.First(&curComment, id)
	if result.Error != nil {
		return result.Error
	}

	if curComment.UserID != userId {
		return errors.New("Invalid Permission")
	}

	result = c.Db.Unscoped().Delete(&models.Comment{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Update implements CommentRepository.
func (c *commentRepositoryImp) Update(id interface{}, userId interface{}, comment models.Comment) error {

	var curComment models.Comment
	result := c.Db.First(&curComment, id)
	if result.Error != nil {
		return result.Error
	}

	if curComment.UserID != userId {
		return errors.New("Invalid Permission")
	}

	curComment.Content = comment.Content

	result = c.Db.Save(&curComment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func NewCommentRepository(Db *gorm.DB) CommentRepository {
	return &commentRepositoryImp{Db: Db}
}
