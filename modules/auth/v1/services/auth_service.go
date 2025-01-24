package service

import repository "github.com/deanfirdianza/dauth-be-go/modules/auth/v1/repositories"

type AuthService interface {
	Login(username, password string) (string, error)
	Register(username, password string) error
}

type authService struct {
	authRepo repository.AuthRepository
}

func NewAuthService(authRepo repository.AuthRepository) AuthService {
	return &authService{authRepo: authRepo}
}

func (s *authService) Login(username, password string) (string, error) {
	// ...implement login logic...
	return "token", nil
}

func (s *authService) Register(username, password string) error {
	// ...implement register logic...
	return nil
}
