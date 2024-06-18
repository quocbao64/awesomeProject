package service

import (
	"awesomeProject/app/constant"
	"awesomeProject/app/domain/dao"
	"awesomeProject/app/domain/dto"
	"awesomeProject/app/pkg"
	"awesomeProject/app/repository"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type CustomerService interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	GetByEmail(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type CustomerServiceImpl struct {
	customerRepo repository.CustomerRepository
}

func (cus CustomerServiceImpl) GetAll(c *gin.Context) {
	data, err := cus.customerRepo.GetAll()

	if err != nil {
		return
	}

	var customerResponse []dto.CustomerResponse
	for _, customer := range data {
		customerResponse = append(customerResponse, daoToDTO(customer))
	}
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), customerResponse))
}

func (cus CustomerServiceImpl) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := cus.customerRepo.GetByID(id)

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), daoToDTO(data)))
}

func (cus CustomerServiceImpl) GetByEmail(c *gin.Context) {
	email := c.Param("email")
	data, err := cus.customerRepo.GetByEmail(email)

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), daoToDTO(data)))
}

func (cus CustomerServiceImpl) Create(c *gin.Context) {
	var customer dao.Customer
	_ = c.BindJSON(&customer)

	passwordHashed, _ := pkg.HashPassword(customer.Password)

	customer.Password = passwordHashed
	customerResponse, err := cus.customerRepo.Create(customer)

	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.InvalidRequest, "Email already exist", pkg.Null()))
			return
		}
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), daoToDTO(customerResponse)))
}

func (cus CustomerServiceImpl) Update(c *gin.Context) {
	var customer dao.Customer
	_ = c.BindJSON(&customer)

	customerResponse, err := cus.customerRepo.Update(customer)

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), daoToDTO(customerResponse)))
}

func (cus CustomerServiceImpl) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := cus.customerRepo.Delete(id)

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), pkg.Null()))
}

func CustomerServiceInit(customerRepo repository.CustomerRepository) *CustomerServiceImpl {
	return &CustomerServiceImpl{
		customerRepo: customerRepo,
	}
}

func daoToDTO(customer dao.Customer) dto.CustomerResponse {
	return dto.CustomerResponse{
		ID:          customer.ID,
		Name:        customer.Name,
		Address:     customer.Address,
		Email:       customer.Email,
		PhoneNumber: customer.PhoneNumber,
		CreatedAt:   customer.CreatedAt,
		UpdatedAt:   customer.UpdatedAt,
		DeletedAt:   customer.DeletedAt,
	}
}
