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

type EditPostRequest struct {
	Title    string `json:"title"`
	Contenct string `json:"content"`
}
type EditPostResponse struct {
	PostId         uuid.UUID        `json:"post_id"`
	AuthorId       uuid.UUID        `json:"author_id"`
	IdempotencyKey string           `json:"indempotency_key"`
	Title          string           `json:"title"`
	Content        string           `json:"content"`
	Status         types.PostStatus `json:"status"`
	CreatedAt      time.Time        `json:"created_at"`
	UpdatedAt      time.Time        `json:"updated_at"`
}

type PublishPostRequest struct {
	Status types.PostStatus `json:"status"`
}

type PublishPostResponse struct {
	PostId uuid.UUID `json:"post_id"`
}
