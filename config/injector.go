//go:build wireinject
// +build wireinject

package config

import (
	"awesomeProject/internal/app/controller"
	"awesomeProject/internal/app/repository"
	"awesomeProject/internal/app/service"
	"github.com/google/wire"
)

var db = wire.NewSet(ConnectDB)

var customerSvcSet = wire.NewSet(service.CustomerServiceInit,
	wire.Bind(new(service.CustomerService), new(*service.CustomerServiceImpl)))

var customerRepoSet = wire.NewSet(repository.CustomerRepositoryInit,
	wire.Bind(new(repository.CustomerRepository), new(*repository.CustomerRepositoryImpl)))

var customerCtrlSet = wire.NewSet(controller.CustomerControllerInit,
	wire.Bind(new(controller.CustomerController), new(*controller.CustomerControllerImpl)))

var authSvcSet = wire.NewSet(service.AuthServiceInit,
	wire.Bind(new(service.AuthService), new(*service.AuthServiceImpl)))

var authCtrlSet = wire.NewSet(controller.AuthControllerInit,
	wire.Bind(new(controller.AuthController), new(*controller.AuthControllerImpl)))

func Init() *Initialize {
	wire.Build(db, customerSvcSet, customerRepoSet, customerCtrlSet, authCtrlSet, authSvcSet, NewInitialize)
	return nil
}
