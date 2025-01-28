package service

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	repository "github.com/deanfirdianza/dauth-be-go/modules/auth/v1/repositories"
	"github.com/deanfirdianza/dauth-be-go/modules/users/v1/models"
	userRepo "github.com/deanfirdianza/dauth-be-go/modules/users/v1/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(username, password string) (models.Accounts, error)
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

func (s *authService) Login(username, password string) (models.Accounts, error) {
	// ...implement login logic...
	userModel, err := s.userRepo.SelectUser(username)
	if err != nil {
		return models.Accounts{}, err
	}

	// Check if the password is correct
	err = bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(password+userModel.Salt))
	if err != nil {
		return models.Accounts{}, fmt.Errorf("invalid password")
	}

	return *userModel, nil
}

func (s *authService) Register(username, email, password string) error {
	// ...implement register logic...

	salt, err := generateRandomSalt()
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
	// Hash the password + salt using bcrypt
	fmt.Println("password : ", password+salt)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func generateRandomSalt() (string, error) {
	salt := make([]byte, 16) // 16 bytes of random data
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	fmt.Println("salt : ", base64.StdEncoding.EncodeToString(salt))
	return base64.StdEncoding.EncodeToString(salt), nil
}
