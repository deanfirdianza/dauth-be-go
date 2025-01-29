package driver

import (
	handler "github.com/deanfirdianza/dauth-be-go/modules/auths/v1/handlers"
	repository "github.com/deanfirdianza/dauth-be-go/modules/auths/v1/repositories"
	service "github.com/deanfirdianza/dauth-be-go/modules/auths/v1/services"
)

var (
	AuthRepo    = repository.NewAuthRepository(DBSqlx)
	AuthService = service.NewAuthService(Conf.App.Secret_key, AuthRepo, UserRepo)
	AuthHandler = handler.NewAuthHandler(AuthService)
)
