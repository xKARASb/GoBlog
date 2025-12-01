package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/xkarasb/blog/pkg/types"
)

type PostDB struct {
	PostId         uuid.UUID        `json:"post_id" db:"post_id"`
	AuthorId       uuid.UUID        `json:"author_id" db:"author_id"`
	IdempotencyKey string           `json:"indempotency_key" db:"idempotency_key"`
	Title          string           `json:"title" db:"title"`
	Content        string           `json:"content" db:"content"`
	CreatedAt      time.Time        `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time        `json:"updated_at" db:"updated_at"`
	Status         types.PostStatus `json:"status" db:"status"`
}

type CreatePostRequest struct {
	IdempotencyKey string `json:"idempotency_key"`
	Title          string `json:"title"`
	Content        string `json:"content"`
}

type CreatePostResponse struct {
	PostId uuid.UUID `json:"post_id"`
}
