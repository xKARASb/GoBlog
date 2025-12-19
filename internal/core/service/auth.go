package service

import (
	"regexp"
	"time"

	"github.com/google/uuid"
	"github.com/xkarasb/blog/internal/core/dto"
	"github.com/xkarasb/blog/pkg/errors"
	"github.com/xkarasb/blog/pkg/hash"
	"github.com/xkarasb/blog/pkg/jwt"
)

type AuthRepository interface {
	AddNewUser(email, password_hash, role, refreshToken string) (*dto.UserDB, error)
	GetUserByEmail(email string) (*dto.UserDB, error)
	GetUserById(id uuid.UUID) (*dto.UserDB, error)
	UpdateRefreshToken(id uuid.UUID, refreshToken string) (*dto.UserDB, error)
}

type AuthService struct {
	rep    AuthRepository
	secret string
}

func NewAuthService(rep AuthRepository, secret string) *AuthService {
	return &AuthService{
		rep,
		secret,
	}
}

func (s *AuthService) validateEmail(email string) bool {
	basicPattern := `^[^@]+@[^@]+\.[^@]+$`
	basicRegex := regexp.MustCompile(basicPattern)
	return basicRegex.MatchString(email)
}

func (s *AuthService) validatePassword(source, db string) bool {
	res, err := hash.CheckPasswordHash(source, db)
	if err != nil {
		return false
	}
	return res
}

func (s *AuthService) RegistrateUser(user *dto.RegistrateUserRequest) (*dto.RegistrateUserResponse, error) {
	if !s.validateEmail(user.Email) {
		return nil, errors.ErrorServiceEmailInvalid
	}

	passwordHash, err := hash.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	refreshToken, err := jwt.NewRefreshToken(user.Email, s.secret)
	if err != nil {
		return nil, err
	}
	newUser, err := s.rep.AddNewUser(user.Email, passwordHash, string(user.Role), refreshToken)

	if err != nil {
		return nil, err
	}
	accessToken := jwt.NewAccessToken(newUser.UserId, s.secret, time.Duration(time.Hour*2))

	resUser := &dto.RegistrateUserResponse{
		Id:           newUser.UserId,
		AccessToken:  accessToken,
		RefreshToken: newUser.RefreshToken,
	}

	return resUser, nil
}

func (s *AuthService) LoginUser(user *dto.LoginUserRequest) (*dto.LoginUserResponse, error) {
	dbUser, err := s.rep.GetUserByEmail(user.Email)

	if err != nil {
		return nil, errors.ErrorRepositoryEmailNotExsist
	}

	if !s.validatePassword(user.Password, dbUser.PasswordHash) {
		return nil, errors.ErrorRepositoryEmailNotExsist
	}

	refreshToken, err := jwt.NewRefreshToken(dbUser.Email, s.secret)
	if err != nil {
		return nil, err
	}

	dbUser, err = s.rep.UpdateRefreshToken(dbUser.UserId, refreshToken)
	if err != nil {
		return nil, errors.ErrorRepositoryEmailNotExsist
	}

	accessToken := jwt.NewAccessToken(dbUser.UserId, s.secret, time.Duration(time.Hour*2))

	resUser := &dto.LoginUserResponse{
		Id:           dbUser.UserId,
		AccessToken:  accessToken,
		RefreshToken: dbUser.RefreshToken,
	}
	return resUser, nil
}

func (s *AuthService) RefreshToken(token *dto.RefreshRequest) (*dto.RefreshResponse, error) {
	claims, err := jwt.ValidateToken(token.RefreshToken, s.secret)

	if err != nil {
		return nil, err
	}

	email, ok := (*claims)["sub"].(string)
	if !ok {
		return nil, errors.ErrorInvalidToken
	}
	dbUser, err := s.rep.GetUserByEmail(email)

	if err != nil {
		return nil, err
	}

	if dbUser.RefreshToken != token.RefreshToken {
		return nil, errors.ErrorInvalidToken
	}

	accessToken := jwt.NewAccessToken(dbUser.UserId, s.secret, time.Duration(time.Hour*2))

	return &dto.RefreshResponse{AccessToken: accessToken}, nil
}

func (s *AuthService) AuthorizeUser(token string) (*dto.UserDB, error) {
	claims, err := jwt.ValidateToken(token, s.secret)
	if err != nil {
		return nil, err
	}

	raw, ok := (*claims)["sub"].(string)
	if !ok {
		return nil, errors.ErrorInvalidToken
	}
	id, err := uuid.Parse(raw)
	if err != nil {
		return nil, errors.ErrorInvalidToken
	}
	data, err := s.rep.GetUserById(id)

	if err != nil {
		return nil, err
	}
	return data, nil
}
