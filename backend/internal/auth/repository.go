package auth

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(user *User) error
	FindByEmail(email string) (*User, error)
	GetMe(userID uuid.UUID) (*User, error)
	GetAll() ([]User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *repository) FindByEmail(email string) (*User, error) {
	var user User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *repository) GetAll() ([]User, error) {
	var users []User

	err := r.db.Order("created_at DESC").Find(&users).Error

	return users, err
}

func (r *repository) GetMe(userID uuid.UUID) (*User, error) {
	var user User

	err := r.db.Where("id = ?", userID).First(&user).Error

	return &user, err
}
