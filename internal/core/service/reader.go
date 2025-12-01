package service

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/xkarasb/blog/internal/core/dto"
	"github.com/xkarasb/blog/pkg/errors"
)

type ReaderRepository interface {
	GetPostByIdempotencyKey(idempotencyKey string) (*dto.PostDB, error)
	CreatePost(
		authorId uuid.UUID,
		idempotencyKey string,
		title,
		content string,
	) (*dto.PostDB, error)
}

type ReaderService struct {
	rep ReaderRepository
}

func NewReaderService(rep ReaderRepository) *ReaderService {
	return &ReaderService{
		rep,
	}
}

func (s *ReaderService) NewPost(authorId uuid.UUID, post *dto.CreatePostRequest) (*dto.CreatePostResponse, error) {
	dbPost, err := s.rep.GetPostByIdempotencyKey(post.IdempotencyKey)

	if dbPost != nil {
		return nil, errors.ErrorKeyIdempotencyAlreadyUsed
	}

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	dbPost, err = s.rep.CreatePost(
		authorId,
		post.IdempotencyKey,
		post.Title,
		post.Content,
	)

	if err != nil {
		return nil, err
	}

	resPost := &dto.CreatePostResponse{
		PostId: dbPost.AuthorId,
	}

	return resPost, nil

}
