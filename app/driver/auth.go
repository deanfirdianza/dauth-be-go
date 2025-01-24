package driver

import (
	handler "github.com/deanfirdianza/dauth-be-go/modules/auth/v1/handlers"
	repository "github.com/deanfirdianza/dauth-be-go/modules/auth/v1/repositories"
	service "github.com/deanfirdianza/dauth-be-go/modules/auth/v1/services"
)

var (
	AuthRepo    = repository.NewAuthRepository(DBSqlx)
	AuthService = service.NewAuthService(AuthRepo)
	AuthHandler = handler.NewAuthHandler(AuthService)
)
