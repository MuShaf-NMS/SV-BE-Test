package repository

import (
	"github.com/MuShaf-NMS/SV-BE-Test/entity"
	"gorm.io/gorm"
)

type PostRepository interface {
	GetAll(limit, offset int, status string) ([]entity.Post, error)
	Create(post *entity.Post) error
	GetOne(id uint) (entity.Post, error)
	Update(post *entity.Post, id uint) error
	Delete(id uint) error
}

type postRepository struct {
	db *gorm.DB
}

func (pr *postRepository) GetAll(limit, offset int, status string) ([]entity.Post, error) {
	var posts []entity.Post
	query := pr.db.Offset(offset)
	if limit > 0 {
		query = query.Limit(limit)
	}
	if status != "" {
		query.Where(&entity.Post{Status: status})
	}
	err := query.Order("id").Find(&posts).Error
	return posts, err
}

func (pr *postRepository) Create(post *entity.Post) error {
	err := pr.db.Create(post).Error
	return err
}

func (pr *postRepository) GetOne(id uint) (entity.Post, error) {
	var post entity.Post
	err := pr.db.Where(&entity.Post{ID: id}).First(&post).Error
	return post, err
}

func (pr *postRepository) Update(post *entity.Post, id uint) error {
	err := pr.db.Where(&entity.Post{ID: id}).Updates(post).Error
	return err
}

func (pr *postRepository) Delete(id uint) error {
	err := pr.db.Where(&entity.Post{ID: id}).Delete(entity.Post{}).Error
	return err
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{
		db: db,
	}
}
