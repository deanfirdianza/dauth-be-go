package services

import (
	"fmt"

	"github.com/deanfirdianza/dauth-be-go/modules/users/v1/models"
	"github.com/deanfirdianza/dauth-be-go/modules/users/v1/repositories"
)

type UserService interface {
	GetUserDetail(userID string) (*models.Accounts, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) GetUserDetail(userID string) (*models.Accounts, error) {
	fmt.Println("userID :", userID)
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}
	fmt.Println("user :", user)
	return user, nil
}
