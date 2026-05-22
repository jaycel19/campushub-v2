package comment

import (
	"errors"

	"github.com/google/uuid"
)

type Service interface {
	CreateComment(comment *Comment) error
	GetComments(postID uuid.UUID) ([]Comment, error)
	DeleteComment(commentID, userID uuid.UUID) error
}

type service struct {
	repo Repository
}

var ErrUnauthorized = errors.New("unauthorized")

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) CreateComment(comment *Comment) error {
	return s.repo.Create(comment)
}

func (s *service) GetComments(postID uuid.UUID) ([]Comment, error) {
	return s.repo.GetByPostID(postID)
}

func (s *service) DeleteComment(commentID, userID uuid.UUID) error {

	comment, err := s.repo.GetByID(commentID)
	if err != nil {
		return err
	}

	if comment.UserID != userID {
		return ErrUnauthorized
	}

	return s.repo.Delete(commentID)
}
