package dto

import (
	"time"

	"github.com/google/uuid"
)

type PostDB struct {
	PostId         uuid.UUID `json:"post_id" db:"post_id"`
	AuthorID       uuid.UUID `json:"author_id" db:"author_id"`
	IdempotencyKey string    `json:"indempotency_key" db:"indempotency_key"`
	Title          string    `json:"title" db:"title"`
	Content        string    `json:"content" db:"content"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	Status         string    `json:"status" db:"status"`
	Images         []ImageDB `json:"images" db:"images"`
}
