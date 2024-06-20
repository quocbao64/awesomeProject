package controller

import (
	"awesomeProject/internal/app/service"
	"github.com/gin-gonic/gin"
)

type CustomerController interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	GetByEmail(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type CustomerControllerImpl struct {
	customerSvc service.CustomerService
}

// GetAll godoc
// @Summary Get all customers
// @Description Get all customers
// @Tags customers
// @Accept json
// @Produce json
// @Success 200 {object} dto.CustomerResponse
// @Router /customers [get]
func (cus CustomerControllerImpl) GetAll(c *gin.Context) {

	cus.customerSvc.GetAll(c)
}

// GetByID GetAll godoc
// @Summary Get customer by ID
// @Description Get customer by ID
// @Tags customers
// @Accept json
// @Produce json
// @Param id path int true "Customer ID"
// @Success 200 {object} dto.CustomerResponse
// @Router /customers/{id} [get]
func (cus CustomerControllerImpl) GetByID(c *gin.Context) {
	cus.customerSvc.GetByID(c)
}

// GetByEmail GetAll godoc
// @Summary Get customer by email
// @Description Get customer by email
// @Tags customers
// @Accept json
// @Produce json
// @Param email query string false "Customer Email"
// @Success 200 {object} dto.CustomerResponse
// @Router /customers [get]
func (cus CustomerControllerImpl) GetByEmail(c *gin.Context) {
	cus.customerSvc.GetByEmail(c)
}

// Create godoc
// @Summary Create a new customer
// @Description Create a new customer
// @Tags customers
// @Accept json
// @Produce json
// @Param customer body dao.Customer true "Customer Body"
// @Success 200 {object} dto.CustomerResponse
// @Router /customers [post]
func (cus CustomerControllerImpl) Create(c *gin.Context) {
	cus.customerSvc.Create(c)
}

// Update godoc
// @Summary Update a customer
// @Description Update a customer
// @Tags customers
// @Accept json
// @Produce json
// @Param id path int true "Customer ID"
// @Param customer body dao.Customer true "Customer Body"
// @Success 200 {object} dto.CustomerResponse
// @Router /customers/{id} [put]
func (cus CustomerControllerImpl) Update(c *gin.Context) {
	cus.customerSvc.Update(c)
}

// Delete godoc
// @Summary Delete a customer
// @Description Delete a customer
// @Tags customers
// @Accept json
// @Produce json
// @Param id path int true "Customer ID"
// @Success 200 {object} dto.CustomerResponse
// @Router /customers/{id} [delete]
func (cus CustomerControllerImpl) Delete(c *gin.Context) {
	cus.customerSvc.Delete(c)
}

func CustomerControllerInit(customerService service.CustomerService) *CustomerControllerImpl {
	return &CustomerControllerImpl{
		customerSvc: customerService,
	}
}
