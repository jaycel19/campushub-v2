package post

import "gorm.io/gorm"

type Repository interface {
	GetAll() ([]Post, error)
	Create(post *Post) error
	GetPaginated(offset, limit int) ([]Post, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Post, error) {
	var posts []Post
	err := r.db.Find(&posts).Error
	return posts, err
}

func (r *repository) Create(post *Post) error {
	return r.db.Create(post).Error
}

func (r *repository) GetPaginated(offset, limit int) ([]Post, error) {
	var posts []Post

	err := r.db.
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&posts).Error

	return posts, err
}
