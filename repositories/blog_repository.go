package repositories

import (
	"errors"
	"main/models"
	"main/utils"

	"gorm.io/gorm"
)

type BlogRepository interface {
	Create(blog models.BlogRequest) models.Blog
	Update(id interface{}, blog models.BlogRequest) error
	Delete(id interface{}) error
	Find(id interface{}) (models.BlogResponse, error)
	FindAll() ([]models.BlogResponse, error)
}

type BlogRepositoryImp struct {
	Db *gorm.DB
}

func getNewBlog(b *BlogRepositoryImp, blog models.BlogRequest) models.Blog {
	var allTopics []models.Topic
	b.Db.Model(&models.Topic{}).Find(&allTopics)
	var existTopics map[string]bool
	existTopics = make(map[string]bool)
	for _, val := range allTopics {
		existTopics[val.Name] = true
	}

	var topics []models.Topic

	for _, val := range blog.Topics {
		var topic models.Topic
		if existTopics[val] != true {
			topic.Name = val
			b.Db.Create(&topic)
		} else {
			b.Db.Where("name = ? ", val).First(&topic)
		}
		topics = append(topics, topic)
	}

	newBlog := models.Blog{
		Title:      blog.Title,
		Content:    blog.Content,
		ImageCover: blog.ImageCover,
		AuthorID:   blog.AuthorID,
		Topics:     topics,
	}
	return newBlog
}

// CreateBlog implements BlogRepository.
func (b *BlogRepositoryImp) Create(blog models.BlogRequest) models.Blog {

	newBlog := getNewBlog(b, blog)
	result := b.Db.Create(&newBlog)
	if result.Error != nil {
		panic(result.Error)
	}

	return newBlog
}

// UpdateBlog implements BlogRepository.
func (b *BlogRepositoryImp) Update(id interface{}, blog models.BlogRequest) error {

	// Get the blog
	var curBlog models.Blog
	result := b.Db.First(&curBlog, id)

	if result.Error != nil {
		return errors.New("This id not found")
	}

	// Get updated blog
	updatedBlog := getNewBlog(b, blog)

	// Delete the relationship between the blog and the topics that no longer related to this blog
	var existTopics map[uint]bool
	existTopics = make(map[uint]bool)
	for _, val := range updatedBlog.Topics {
		existTopics[val.ID] = true
	}
	var blogs models.Blog
	b.Db.Preload("Topics").First(&blogs, id)
	for _, val := range blogs.Topics {
		if existTopics[val.ID] != true {
			b.Db.Model(&blogs).Association("Topics").Delete(&val)
		}
	}

	// Update the blog
	curBlog.Title = updatedBlog.Title
	curBlog.Content = updatedBlog.Title
	curBlog.ImageCover = updatedBlog.Content
	curBlog.Topics = updatedBlog.Topics

	result = b.Db.Save(&curBlog)
	b.Db.Model(&curBlog).Association("Tags").Replace(curBlog.Topics)
	if result.Error != nil {
		panic(result.Error)
	}

	return nil
}

// DeleteBlog implements BlogRepository.
func (b *BlogRepositoryImp) Delete(id interface{}) error {
	result := b.Db.Unscoped().Delete(&models.Blog{}, id)
	if result.Error != nil {
		panic(result.Error)
	}

	if result.RowsAffected == 0 {
		return errors.New("This blog not found")
	}

	return nil
}

// FindBlog implements BlogRepository.
func (b *BlogRepositoryImp) Find(id interface{}) (models.BlogResponse, error) {
	var blog models.Blog
	result := b.Db.Preload("Author").Preload("Topics").Preload("Comments.User").Preload("Likes.User").First(&blog, id)

	if result.Error != nil {
		return models.BlogResponse{}, result.Error
	}
	blogRes := utils.GetBlogRes(blog)
	return blogRes, nil
}

// FindAll implements BlogRepository.
func (b *BlogRepositoryImp) FindAll() ([]models.BlogResponse, error) {
	var blogs []models.Blog
	result := b.Db.Preload("Author").Preload("Topics").Find(&blogs)
	if result.Error != nil {
		return nil, result.Error
	}

	var blogsRes []models.BlogResponse
	for _, val := range blogs {
		blogsRes = append(blogsRes, utils.GetBlogRes(val))
	}

	return blogsRes, nil
}

func NewBlogRepository(Db *gorm.DB) BlogRepository {
	return &BlogRepositoryImp{Db: Db}
}
