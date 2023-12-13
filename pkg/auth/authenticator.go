package auth

type Strategy interface {
	Authenticate(username, password string) (bool, error)
}

type Authenticator struct {
	Strategy Strategy
}

func (a *Authenticator) Authenticate(username, password string) (bool, error) {
	return a.Strategy.Authenticate(username, password)
}
