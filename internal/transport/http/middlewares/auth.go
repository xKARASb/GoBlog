package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/xkarasb/blog/internal/core/dto"
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
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, "No authorization provided")
			return
		}

		rawToken := strings.Split(auth_header, " ")
		if len(rawToken) != 2 {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, "No authorization provided")
			return
		}
		token := rawToken[1]
		user, err := m.service.AuthorizeUser(token)

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, "No authorization provided")
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
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintln(w, "Incorrect user")
			return
		}
		if user.Role == types.Author {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintln(w, "Incorrect user")
			return
		}
	})
}
