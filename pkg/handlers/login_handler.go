package handlers

import (
	"encoding/json"
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
	// Initialize the structure to store the data
	creds := &struct {
		username string `json:"username"`
		password string `json:"password"`
	}{}

	// Decode the JSON body of the request into the creds structure
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		// Handle error
		return
	}

	ok, err := h.authenticator.Authenticate(creds.username, creds.password)
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
