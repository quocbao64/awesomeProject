package config

import (
	"awesomeProject/app/controller"
	"awesomeProject/app/repository"
	"awesomeProject/app/service"
)

type Initialize struct {
	customerRepo repository.CustomerRepository
	customerSvc  service.CustomerService
	CustomerCtrl controller.CustomerController
	AuthSvc      service.AuthService
	AuthCtrl     controller.AuthController
}

func NewInitialize(
	customerRepo repository.CustomerRepository,
	customerSvc service.CustomerService,
	customerCtrl controller.CustomerController,
	authCtrl controller.AuthController,
	authSvc service.AuthService,
) *Initialize {
	return &Initialize{
		customerRepo: customerRepo,
		customerSvc:  customerSvc,
		CustomerCtrl: customerCtrl,
		AuthCtrl:     authCtrl,
		AuthSvc:      authSvc,
	}
}
