package service

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"

	repository "github.com/deanfirdianza/dauth-be-go/modules/auths/v1/repositories"
	"github.com/deanfirdianza/dauth-be-go/modules/users/v1/models"
	userRepo "github.com/deanfirdianza/dauth-be-go/modules/users/v1/repositories"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	secretKey string
	authRepo  repository.AuthRepository
	userRepo  userRepo.UserRepository
}

type AuthService interface {
	Login(username, password string) (models.Tokens, error)
	Register(username, email, password string) error
	GenerateJWT(username string, expiry time.Duration) (string, error)
	ValidateJWT(tokenString string) (jwt.MapClaims, error)
	RefreshJWTToken(refreshToken string) (models.Tokens, error)
}

func NewAuthService(
	secretKey string,
	authRepo repository.AuthRepository,
	userRepo userRepo.UserRepository,
) AuthService {
	return &authService{
		secretKey: secretKey,
		authRepo:  authRepo,
		userRepo:  userRepo,
	}
}

func (s *authService) Login(username, password string) (models.Tokens, error) {
	// ...implement login logic...
	userModel, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return models.Tokens{}, err
	}

	// Check if the password is correct
	err = bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(password+userModel.Salt))
	if err != nil {
		return models.Tokens{}, fmt.Errorf("invalid password")
	}

	accessToken, err := s.GenerateJWT(userModel.ID.String(), 15*time.Minute) // Access token
	if err != nil {
		return models.Tokens{}, err
	}

	refreshToken, err := s.GenerateJWT(userModel.ID.String(), 7*24*time.Hour) // Refresh token
	if err != nil {
		return models.Tokens{}, err
	}

	// Delete old refresh tokens
	_, err = s.authRepo.DeleteOldAuths(userModel.ID.String())
	if err != nil {
		return models.Tokens{}, err
	}

	_, err = s.authRepo.InsertAuth(userModel.ID.String(), refreshToken, time.Now().Add(7*24*time.Hour), false)
	if err != nil {
		return models.Tokens{}, err
	}

	return models.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *authService) RefreshJWTToken(refreshToken string) (models.Tokens, error) {
	claims, err := s.ValidateJWT(refreshToken)
	if err != nil {
		return models.Tokens{}, fmt.Errorf("invalid refresh token: %v", err)
	}

	userID, ok := claims["uid"].(string)
	if !ok {
		return models.Tokens{}, fmt.Errorf("invalid refresh token claims")
	}

	// Generate new access token
	accessToken, err := s.GenerateJWT(userID, 15*time.Minute)
	if err != nil {
		return models.Tokens{}, err
	}

	// Generate new refresh token
	newRefreshToken, err := s.GenerateJWT(userID, 7*24*time.Hour)
	if err != nil {
		return models.Tokens{}, err
	}

	// Revoke the latest refresh token
	_, err = s.authRepo.RevokeAuth(userID)
	if err != nil {
		return models.Tokens{}, err
	}

	// Delete old refresh tokens
	_, err = s.authRepo.DeleteOldAuths(userID)
	if err != nil {
		return models.Tokens{}, err
	}

	// Insert the new refresh token
	_, err = s.authRepo.InsertAuth(userID, newRefreshToken, time.Now().Add(7*24*time.Hour), false)
	if err != nil {
		return models.Tokens{}, err
	}

	return models.Tokens{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}

func (s *authService) ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}

func (s *authService) GenerateJWT(id string, expiry time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"uid": id,
		"exp": time.Now().Add(expiry).Unix(), // Set expiration time
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(s.secretKey))
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
