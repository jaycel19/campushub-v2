package comment

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jaycel19/campushub/backend/internal/shared"
)

type Service interface {
	CreateComment(
		userIDStr,
		postIDStr,
		content string) (*Comment, error)
	GetComments(postIDStr string) ([]Comment, error)
	DeleteComment(commentIDStr, userIDStr string) error
}

type service struct {
	repo Repository
}

var ErrUnauthorized = errors.New("unauthorized")

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) CreateComment(
	userIDStr,
	postIDStr,
	content string) (*Comment, error) {

	postID, err := uuid.Parse(postIDStr)
	if err != nil {
		return nil, shared.ErrInvalidPostID
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, shared.ErrInvalidUserID
	}
	comment := &Comment{
		PostID:  postID,
		UserID:  userID,
		Content: content,
	}

	err = s.repo.Create(comment)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (s *service) GetComments(postIDStr string) ([]Comment, error) {
	postID, err := uuid.Parse(postIDStr)
	if err != nil {
		return nil, shared.ErrInvalidPostID
	}
	return s.repo.GetByPostID(postID)
}

func (s *service) DeleteComment(commentIDStr, userIDStr string) error {

	commentID, err := uuid.Parse(commentIDStr)
	if err != nil {
		return shared.ErrInvalidCommentID
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return shared.ErrInvalidUserID
	}

	comment, err := s.repo.GetByID(commentID)
	if err != nil {
		return err
	}

	if comment.UserID != userID {
		return ErrUnauthorized
	}

	return s.repo.Delete(commentID)
}
