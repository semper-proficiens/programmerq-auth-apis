package auth

import (
	"fmt"
	"gopkg.in/ldap.v2"
)

type LdapStrategy struct{}

func (l *LdapStrategy) Authenticate(username, password string) (bool, error) {
	conn, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", "ldap.example.com", 389))
	if err != nil {
		return false, err
	}
	defer conn.Close()

	err = conn.Bind(username, password)
	if err != nil {
		return false, err
	}
	return true, nil
}
