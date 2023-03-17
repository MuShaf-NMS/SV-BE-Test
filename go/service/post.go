package service

import (
	"github.com/MuShaf-NMS/SV-BE-Test/dto"
	"github.com/MuShaf-NMS/SV-BE-Test/entity"
	"github.com/MuShaf-NMS/SV-BE-Test/helper"
	"github.com/MuShaf-NMS/SV-BE-Test/repository"
)

type PostService interface {
	GetAll(limit, offset int, status string) ([]entity.Post, error)
	Create(postDto dto.PostDTO) error
	GetOne(id uint) (entity.Post, error)
	Update(postDto dto.PostDTO, id uint) error
	Delete(id uint) error
}

type postService struct {
	repository repository.PostRepository
}

func (pr *postService) GetAll(limit, offset int, status string) ([]entity.Post, error) {
	posts, err := pr.repository.GetAll(limit, offset, status)
	return posts, err
}

func (pr *postService) Create(postDto dto.PostDTO) error {
	post := entity.Post{
		Title:    postDto.Title,
		Content:  postDto.Content,
		Category: postDto.Category,
		Status:   postDto.Status,
	}
	err := pr.repository.Create(&post)
	if err != nil {
		return helper.NewError(422, "Failed to create new post", nil)
	}
	return nil
}

func (pr *postService) GetOne(id uint) (entity.Post, error) {
	post, err := pr.repository.GetOne(id)
	if err != nil {
		return post, helper.NewError(404, "Post not found", nil)
	}
	return post, nil
}

func (pr *postService) Update(postDto dto.PostDTO, id uint) error {
	post := entity.Post{
		Title:    postDto.Title,
		Content:  postDto.Content,
		Category: postDto.Category,
		Status:   postDto.Status,
	}
	err := pr.repository.Update(&post, id)
	if err != nil {
		return helper.NewError(422, "Failed to create new post", nil)
	}
	return nil
}

func (pr *postService) Delete(id uint) error {
	err := pr.repository.Delete(id)
	if err != nil {
		return helper.NewError(404, "Post not found", nil)
	}
	return nil
}

func NewPostService(repository repository.PostRepository) PostService {
	return &postService{
		repository: repository,
	}
}
