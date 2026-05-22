package comment

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	PostID    uuid.UUID `json:"post_id" gorm:"type:uuid;not null"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
