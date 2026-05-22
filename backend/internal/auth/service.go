package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/jaycel19/campushub/backend/internal/shared"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(req *RegisterRequest) error
	Login(email, password string) (string, error)
	GetMe(userIDStr string) (*User, error)
	GetAll() ([]User, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Register(req *RegisterRequest) error {

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return err
	}
	req.Password = string(hashed)
	user := User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	return s.repo.Create(&user)
}

func (s *service) Login(email, password string) (string, error) {
	const TokenDuration = time.Hour * 24
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Create JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(TokenDuration).Unix(),
	})

	secret := os.Getenv("JWT_SECRET")

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *service) GetMe(userIDStr string) (*User, error) {
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, shared.ErrInvalidUserID
	}
	users, err := s.repo.GetMe(userID)
	return users, err
}

func (s *service) GetAll() ([]User, error) {
	users, err := s.repo.GetAll()
	return users, err
}
