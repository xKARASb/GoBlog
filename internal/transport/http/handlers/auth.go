package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/xkarasb/blog/internal/core/dto"
	"github.com/xkarasb/blog/internal/core/service"
	"github.com/xkarasb/blog/pkg/errors"
)

type AuthController struct {
	service *service.AuthService
}

func NewAuthController(service *service.AuthService) *AuthController {
	return &AuthController{service: service}
}
func (c *AuthController) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	reqUser := &dto.RegistrateUserRequest{}
	if err := json.NewDecoder(r.Body).Decode(reqUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Incorrect body")
		return
	}
	resp, err := c.service.RegistrateUser(reqUser)
	if err != nil {
		if err == errors.ErrorRepositoryUserAlreadyExsist {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, "%s\n", err.Error())
			return
		}
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprintf(w, "%s\n", err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (c *AuthController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	reqUser := &dto.LoginUserRequest{}
	if err := json.NewDecoder(r.Body).Decode(reqUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Incorrect body")
		return
	}
	resp, err := c.service.LoginUser(reqUser)

	if err != nil {
		if err == errors.ErrorRepositoryEmailNotExsist {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, "Email or password incorrect\n")
			return
		}
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprintf(w, "%s\n", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (c *AuthController) RefreshHandler(w http.ResponseWriter, r *http.Request) {
	req := &dto.RefreshRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Incorrect body")
		return
	}
	resp, err := c.service.RefreshToken(req)

	if err != nil {
		if err == errors.ErrorInvalidToken {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "refresh token expired or incorrect\n")
			return
		}
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprintf(w, "%s\n", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
