package shared

import "errors"

var (
	ErrUnauthorized     = errors.New("unauthorized")
	ErrForbidden        = errors.New("forbidden")
	ErrNotFound         = errors.New("resource not found")
	ErrInvalidInput     = errors.New("invalid input")
	ErrInvalidPostID    = errors.New("invalid post id")
	ErrInvalidUserID    = errors.New("invalid user id")
	ErrInvalidCommentID = errors.New("invalid comment id")
)
