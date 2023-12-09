package repositories

import (
	"main/models"
	"main/utils"

	"gorm.io/gorm"
)

type TopicRepository interface {
	Create(topic models.Topic) models.Topic
	FindAll() []models.Topic
	Find(id interface{}) (models.TopicResponse, error)
}

type TopicRepositoryImp struct {
	Db *gorm.DB
}

// Create implements TopicRepository.
func (t *TopicRepositoryImp) Create(topic models.Topic) models.Topic {
	result := t.Db.Create(&topic)
	if result.Error != nil {
		panic(result.Error)
	}
	return topic
}

// Create implements TopicRepository.
func (t *TopicRepositoryImp) FindAll() []models.Topic {
	var topics []models.Topic
	result := t.Db.Find(&topics)
	if result.Error != nil {
		panic(result.Error)
	}

	return topics
}

// Find implements TopicRepository.
func (t *TopicRepositoryImp) Find(id interface{}) (models.TopicResponse, error) {
	var topic models.Topic
	result := t.Db.Preload("Blogs.Topics").Preload("Blogs.Author").First(&topic, id)
	if result.Error != nil {
		return models.TopicResponse{}, result.Error
	}

	var topicBlogs []models.BlogResponse
	for _, val := range topic.Blogs {
		topicBlogs = append(topicBlogs, utils.GetBlogRes(val))
	}

	var topicRes models.TopicResponse
	topicRes.Name = topic.Name
	topicRes.Blogs = topicBlogs

	return topicRes, nil
}

func NewTopicRepository(Db *gorm.DB) TopicRepository {
	return &TopicRepositoryImp{Db: Db}
}
