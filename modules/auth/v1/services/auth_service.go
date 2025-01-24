package service

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	repository "github.com/deanfirdianza/dauth-be-go/modules/auth/v1/repositories"
	userRepo "github.com/deanfirdianza/dauth-be-go/modules/users/v1/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(username, password string) (string, error)
	Register(username, email, password string) error
}

type authService struct {
	authRepo repository.AuthRepository
	userRepo userRepo.UserRepository
}

func NewAuthService(
	authRepo repository.AuthRepository,
	userRepo userRepo.UserRepository,
) AuthService {
	return &authService{
		authRepo: authRepo,
		userRepo: userRepo,
	}
}

func (s *authService) Login(username, password string) (string, error) {
	// ...implement login logic...
	return "token", nil
}

func (s *authService) Register(username, email, password string) error {
	// ...implement register logic...

	salt, err := GenerateRandomSalt()
	if err != nil {
		return err
	}

	hashedPassword, err := generatePasswordHash(password, salt)
	if err != nil {
		return err
	}

	userID, err := s.userRepo.InsertUser(username, hashedPassword, email, salt)
	if err != nil {
		return err
	}
	fmt.Println("userID : ", userID)

	return nil
}

func generatePasswordHash(password string, salt string) (string, error) {
	// Combine password and salt
	passwordWithSalt := password + salt

	// Hash the password + salt using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordWithSalt), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func GenerateRandomSalt() (string, error) {
	salt := make([]byte, 16) // 16 bytes of random data
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}
