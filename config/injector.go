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

var categorySvcSet = wire.NewSet(service.CategoryServiceInit,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)))

var categoryRepoSet = wire.NewSet(repository.CategoryRepositoryInit,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)))

var categoryCtrlSet = wire.NewSet(controller.CategoryControllerInit,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)))

var productSvcSet = wire.NewSet(service.ProductServiceInit,
	wire.Bind(new(service.ProductService), new(*service.ProductServiceImpl)))

var productRepoSet = wire.NewSet(repository.ProductRepositoryInit,
	wire.Bind(new(repository.ProductRepository), new(*repository.ProductRepositoryImpl)))

var productCtrlSet = wire.NewSet(controller.ProductControllerInit,
	wire.Bind(new(controller.ProductController), new(*controller.ProductControllerImpl)))

func Init() *Initialize {
	wire.Build(
		db,
		customerSvcSet,
		customerRepoSet,
		customerCtrlSet,
		authCtrlSet,
		authSvcSet,
		categoryRepoSet,
		categorySvcSet,
		categoryCtrlSet,
		productRepoSet,
		productSvcSet,
		productCtrlSet,
		NewInitialize)
	return nil
}
