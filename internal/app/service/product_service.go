package service

import (
	"awesomeProject/internal/app/constant"
	"awesomeProject/internal/app/domain/dao"
	"awesomeProject/internal/app/pkg"
	"awesomeProject/internal/app/repository"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type ProductService interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type ProductServiceImpl struct {
	productRepo repository.ProductRepository
}

func (p *ProductServiceImpl) GetAll(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	data, err := p.productRepo.GetAll(limit, offset)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (p *ProductServiceImpl) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := p.productRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, pkg.BuildResponse(constant.DataNotFound, err, pkg.Null()))
			return
		}
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (p *ProductServiceImpl) Create(c *gin.Context) {
	var product dao.Product
	_ = c.BindJSON(&product)
	data, err := p.productRepo.Create(product)
	if err != nil {
		if errors.Is(err, gorm.ErrInvalidData) ||
			errors.Is(err, gorm.ErrForeignKeyViolated) {
			c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.InvalidRequest, err, pkg.Null()))
			return
		}
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (p *ProductServiceImpl) Update(c *gin.Context) {
	var product dao.Product
	_ = c.BindJSON(&product)
	data, err := p.productRepo.Update(product)
	if err != nil {
		if errors.Is(err, gorm.ErrInvalidData) ||
			errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.InvalidRequest, err, pkg.Null()))
			return
		}
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (p *ProductServiceImpl) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := p.productRepo.Delete(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, pkg.BuildResponse(constant.DataNotFound, err, pkg.Null()))
			return
		}
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), pkg.Null()))
}

func ProductServiceInit(productRepo repository.ProductRepository) *ProductServiceImpl {
	return &ProductServiceImpl{productRepo: productRepo}
}
