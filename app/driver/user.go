package driver

import (
	handler "github.com/deanfirdianza/dauth-be-go/modules/users/v1/handlers"
	repository "github.com/deanfirdianza/dauth-be-go/modules/users/v1/repositories"
	service "github.com/deanfirdianza/dauth-be-go/modules/users/v1/services"
)

var (
	UserRepo    = repository.NewUserRepository(DBSqlx)
	UserService = service.NewUserService(UserRepo)
	UserHandler = handler.NewUserHandler(UserService)
)
