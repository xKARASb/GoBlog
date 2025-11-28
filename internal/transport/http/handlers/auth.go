package handlers

import (
	"fmt"
	"net/http"
)

type AuthController struct{}

func NewAuthController() *AuthController {
	return &AuthController{}
}
func (c *AuthController) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Auth %s\n", r.URL)
}

func (c *AuthController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Auth %s\n", r.URL)
}

func (c *AuthController) RefreshHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Auth %s\n", r.URL)
}
