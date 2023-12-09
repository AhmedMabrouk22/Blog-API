package services

import (
	"main/models"
	"main/repositories"
)

type BlogServices interface {
	Create(blog models.BlogRequest) models.Blog
	Update(id interface{}, blog models.BlogRequest) error
	Delete(id interface{}) error
	Find(id interface{}) (models.BlogResponse, error)
	FindAll() ([]models.BlogResponse, error)
}

type BlogServicesImp struct {
	blogRepository repositories.BlogRepository
}

// CreateBlog implements BlogServices.
func (b *BlogServicesImp) Create(blog models.BlogRequest) models.Blog {
	newBlog := b.blogRepository.Create(blog)
	return newBlog
}

// UpdateBlog implements BlogServices.
func (b *BlogServicesImp) Update(id interface{}, blog models.BlogRequest) error {
	err := b.blogRepository.Update(id, blog)
	return err
}

// DeleteBlog implements BlogServices.
func (b *BlogServicesImp) Delete(id interface{}) error {
	err := b.blogRepository.Delete(id)
	return err
}

// FindBlog implements BlogServices.
func (b *BlogServicesImp) Find(id interface{}) (models.BlogResponse, error) {
	blog, err := b.blogRepository.Find(id)
	if err != nil {
		return models.BlogResponse{}, err
	}
	return blog, nil
}

// FindAll implements BlogServices.
func (b *BlogServicesImp) FindAll() ([]models.BlogResponse, error) {
	blogs, err := b.blogRepository.FindAll()
	return blogs, err
}

func NewBlogServices(blogRepository repositories.BlogRepository) BlogServices {
	return &BlogServicesImp{blogRepository: blogRepository}
}
