package services

import (
	"github.com/deanfirdianza/dauth-be-go/modules/users/v1/repositories"
)

type UserService interface {
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}
