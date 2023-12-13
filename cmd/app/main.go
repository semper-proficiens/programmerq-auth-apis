package main

import (
	"net/http"
	"programmerq-auth-apis/pkg/auth"
	"programmerq-auth-apis/pkg/handlers"
)

func main() {
	// Initialize strategies
	ldapStrategy := &auth.LdapStrategy{}
	oauth2Strategy := &auth.Oauth2Strategy{}

	// Initialize handlers with the strategies
	ldapHandler := handlers.NewLoginHandler(ldapStrategy)
	oauth2Handler := handlers.NewLoginHandler(oauth2Strategy)

	// Setup routes
	http.HandleFunc("/auth/ldap", ldapHandler.HandleLogin)
	http.HandleFunc("/auth/oauth2", oauth2Handler.HandleLogin)

	// Start the server
	http.ListenAndServe(":8080", nil)
}
