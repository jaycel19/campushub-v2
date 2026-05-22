package comment

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(comment *Comment) error
	GetByPostID(postID uuid.UUID) ([]Comment, error)
	Delete(commentID uuid.UUID) error
	GetByID(commentID uuid.UUID) (*Comment, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(comment *Comment) error {
	return r.db.Create(comment).Error
}

func (r *repository) GetByPostID(postID uuid.UUID) ([]Comment, error) {
	var comments []Comment

	err := r.db.
		Where("post_id = ?", postID).
		Order("created_at ASC").
		Find(&comments).Error

	return comments, err
}

func (r *repository) Delete(commentID uuid.UUID) error {
	return r.db.Delete(&Comment{}, "id = ?", commentID).Error
}

func (r *repository) GetByID(commentID uuid.UUID) (*Comment, error) {
	var comment Comment

	err := r.db.Where("id = ?", commentID).First(&comment).Error

	return &comment, err
}
