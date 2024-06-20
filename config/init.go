package config

import (
	"awesomeProject/internal/app/controller"
	"awesomeProject/internal/app/repository"
	"awesomeProject/internal/app/service"
)

type Initialize struct {
	customerRepo repository.CustomerRepository
	customerSvc  service.CustomerService
	CustomerCtrl controller.CustomerController
	AuthSvc      service.AuthService
	AuthCtrl     controller.AuthController
	categoryRepo repository.CategoryRepository
	categorySvc  service.CategoryService
	CategoryCtrl controller.CategoryController
	productRepo  repository.ProductRepository
	productSvc   service.ProductService
	ProductCtrl  controller.ProductController
}

func NewInitialize(
	customerRepo repository.CustomerRepository,
	customerSvc service.CustomerService,
	customerCtrl controller.CustomerController,
	authCtrl controller.AuthController,
	authSvc service.AuthService,
	categoryRepo repository.CategoryRepository,
	categorySvc service.CategoryService,
	categoryCtrl controller.CategoryController,
	productRepo repository.ProductRepository,
	productSvc service.ProductService,
	productCtrl controller.ProductController,
) *Initialize {
	return &Initialize{
		customerRepo: customerRepo,
		customerSvc:  customerSvc,
		CustomerCtrl: customerCtrl,
		AuthCtrl:     authCtrl,
		AuthSvc:      authSvc,
		categoryRepo: categoryRepo,
		categorySvc:  categorySvc,
		CategoryCtrl: categoryCtrl,
		productRepo:  productRepo,
		productSvc:   productSvc,
		ProductCtrl:  productCtrl,
	}
}
