package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/xkarasb/blog/internal/core/dto"
)

type AuthService interface {
	AuthorizeUser(token string) (*dto.UserDB, error)
}

func AuthMiddleware(userService AuthService, secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
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
			fmt.Println(token)
			user, err := userService.AuthorizeUser(token)

			fmt.Println(user)
			fmt.Println(err)
			if err != nil {
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, "No authorization provided")
				return
			}

			ctx := context.WithValue(r.Context(), "user", user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
