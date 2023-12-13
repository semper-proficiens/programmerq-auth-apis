package handlers

import (
	"net/http"
	"programmerq-auth-apis/pkg/auth"
)

type LoginHandler struct {
	authenticator *auth.Authenticator
}

func NewLoginHandler(strategy auth.Strategy) *LoginHandler {
	return &LoginHandler{
		authenticator: &auth.Authenticator{Strategy: strategy},
	}
}

func (h *LoginHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	// Extract credentials from the request
	// username := ...
	// password := ...

	ok, err := h.authenticator.Authenticate(username, password)
	if err != nil {
		// Handle error
		return
	}

	if !ok {
		// Handle failed login
		return
	}

	// Handle successful login
}
