package services

import (
	"main/models"
	"main/repositories"
)

type TopicServices interface {
	Create(topic models.Topic) models.Topic
	FindAll() []models.Topic
	Find(id interface{}) (models.TopicResponse, error)
}

type TopicServicesImp struct {
	topicRepository repositories.TopicRepository
}

// Create implements TopicServices.
func (t *TopicServicesImp) Create(topic models.Topic) models.Topic {
	result := t.topicRepository.Create(topic)
	return result
}

// Create implements TopicServices.
func (t *TopicServicesImp) FindAll() []models.Topic {
	result := t.topicRepository.FindAll()
	return result
}

// Create implements TopicServices.
func (t *TopicServicesImp) Find(id interface{}) (models.TopicResponse, error) {
	return t.topicRepository.Find(id)
}

func NewTopicServices(topicRepository repositories.TopicRepository) TopicServices {
	return &TopicServicesImp{topicRepository: topicRepository}
}
