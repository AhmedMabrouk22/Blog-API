package repositories

import (
	"errors"
	"main/models"

	"gorm.io/gorm"
)

type LikeRepository interface {
	Add(blogId uint, userId uint) error
	Delete(blogId uint, userId uint) error
	Get(blogId uint) ([]models.LikeRes, error)
}

type likeRepositoryImp struct {
	Db *gorm.DB
}

// Get implements LikeRepository.
func (l *likeRepositoryImp) Get(blogId uint) ([]models.LikeRes, error) {
	var likes []models.Like
	result := l.Db.First(&models.Blog{}, blogId)
	if result.Error != nil {
		return []models.LikeRes{}, result.Error
	}

	result = l.Db.Where("blog_id = ? ", blogId).Preload("User").Find(&likes)

	var likesRes []models.LikeRes
	for _, val := range likes {
		like := models.LikeRes{
			ID: val.ID,
			User: models.LikeUser{
				ID:    val.User.ID,
				Name:  val.User.Name,
				Image: val.User.Image,
			},
		}
		likesRes = append(likesRes, like)
	}
	return likesRes, nil
}

// Add implements LikeRepository.
func (l *likeRepositoryImp) Add(blogId uint, userId uint) error {
	var like models.Like

	result := l.Db.First(&models.Blog{}, blogId)
	if result.Error != nil {
		return result.Error
	}

	result = l.Db.Where("user_id = ? AND blog_id = ?", userId, blogId).First(&like)

	if result.Error == nil {
		return errors.New("Something error")
	}

	newLike := models.Like{
		UserID: userId,
		BlogID: blogId,
	}
	result = l.Db.Create(&newLike)
	return result.Error
}

// Delete implements LikeRepository.
func (l *likeRepositoryImp) Delete(blogId uint, userId uint) error {
	result := l.Db.First(&models.Blog{}, blogId)
	if result.Error != nil {
		return result.Error
	}

	var like models.Like
	result = l.Db.Where("user_id = ? AND blog_id = ?", userId, blogId).First(&like)

	if result.Error != nil {
		return errors.New("Something error")
	}

	result = l.Db.Unscoped().Where("user_id = ? AND blog_id = ?", userId, blogId).Delete(models.Like{})
	return result.Error

}

func NewLikeRepository(Db *gorm.DB) LikeRepository {
	return &likeRepositoryImp{Db: Db}
}
