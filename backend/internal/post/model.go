package post

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    uint
	Content   string
	Likes     int
	Comments  int
	CreatedAt time.Time
	UpdatedAt time.Time
}
