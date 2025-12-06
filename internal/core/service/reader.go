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
	GetPublishedPosts() ([]*dto.PostUserDB, error)
	GetUserPosts(userId uuid.UUID) ([]*dto.PostUserDB, error)
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
		PostId: dbPost.PostId,
	}

	return resPost, nil
}

func (s *ReaderService) GetPublishedPosts() ([]*dto.GetPostResponse, error) {
	posts, err := s.rep.GetPublishedPosts()

	if err != nil {
		return nil, err
	}

	res := make([]*dto.GetPostResponse, len(posts))

	for i, raw := range posts {
		res[i] = &dto.GetPostResponse{
			PostId: raw.PostId,
			Author: dto.UserResponse{
				UserId: raw.AuthorId,
				Email:  raw.Email,
			},
			Title:     raw.Title,
			Content:   raw.Content,
			Status:    raw.Status,
			CreatedAt: raw.CreatedAt,
			UpdatedAt: raw.UpdatedAt,
		}
	}
	return res, nil
}

func (s *ReaderService) GetAuthorPosts(authorId uuid.UUID) ([]*dto.GetPostResponse, error) {
	posts, err := s.rep.GetUserPosts(authorId)

	if err != nil {
		return nil, err
	}

	res := make([]*dto.GetPostResponse, len(posts))

	for i, raw := range posts {
		res[i] = &dto.GetPostResponse{
			PostId: raw.PostId,
			Author: dto.UserResponse{
				UserId: raw.AuthorId,
				Email:  raw.Email,
			},
			Title:     raw.Title,
			Content:   raw.Content,
			Status:    raw.Status,
			CreatedAt: raw.CreatedAt,
			UpdatedAt: raw.UpdatedAt,
		}
	}
	return res, nil
}
