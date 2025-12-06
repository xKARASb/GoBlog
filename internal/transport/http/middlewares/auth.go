package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/xkarasb/blog/internal/core/dto"
	"github.com/xkarasb/blog/pkg/errors"
	"github.com/xkarasb/blog/pkg/types"
)

type AuthService interface {
	AuthorizeUser(token string) (*dto.UserDB, error)
}

type AuthMiddlewareManager struct {
	service AuthService
}

func NewAuthMiddlewareManager(service AuthService) *AuthMiddlewareManager {
	return &AuthMiddlewareManager{service}
}

func (m *AuthMiddlewareManager) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth_header := r.Header.Get("Authorization")
		if auth_header == "" {
			http.Error(w, errors.ErrorHttpNoAuth.Error(), http.StatusForbidden)
			return
		}

		rawToken := strings.Split(auth_header, " ")
		if len(rawToken) != 2 {
			http.Error(w, errors.ErrorHttpNoAuth.Error(), http.StatusForbidden)
			return
		}
		token := rawToken[1]
		user, err := m.service.AuthorizeUser(token)

		if err != nil {
			http.Error(w, errors.ErrorHttpNoAuth.Error(), http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), types.CtxUser, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *AuthMiddlewareManager) AuthorOnlyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		userRaw := ctx.Value(types.CtxUser)
		user, ok := userRaw.(*dto.UserDB)
		if !ok {
			http.Error(w, errors.ErrorHttpIncorrectUser.Error(), http.StatusForbidden)
			return
		}
		if user.Role == types.Author {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, errors.ErrorHttpIncorrectUser.Error(), http.StatusForbidden)
		}
	})
}
