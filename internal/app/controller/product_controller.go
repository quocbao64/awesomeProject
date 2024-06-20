package controller

import (
	"awesomeProject/internal/app/service"
	"github.com/gin-gonic/gin"
)

type ProductController interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type ProductControllerImpl struct {
	productSvc service.ProductService
}

// GetAll godoc
// @Summary Get all products
// @Description Get all products
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {object} dao.Product
// @Router /products [get]
func (p *ProductControllerImpl) GetAll(c *gin.Context) {
	p.productSvc.GetAll(c)
}

// GetByID godoc
// @Summary Get product by ID
// @Description Get product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} dao.Product
// @Router /products/{id} [get]
func (p *ProductControllerImpl) GetByID(c *gin.Context) {
	p.productSvc.GetByID(c)
}

// Create godoc
// @Summary Create product
// @Description Create product
// @Tags products
// @Accept json
// @Produce json
// @Param name body string true "Product Name"
// @Param price body number true "Product Price"
// @Param category_id body int true "Category ID"
// @Success 200 {object} dao.Product
// @Router /products [post]
func (p *ProductControllerImpl) Create(c *gin.Context) {
	p.productSvc.Create(c)
}

// Update godoc
// @Summary Update product
// @Description Update product
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param name body string true "Product Name"
// @Param price body int true "Product Price"
// @Param category_id body int true "Category ID"
// @Success 200 {object} dao.Product
// @Router /products/{id} [put]
func (p *ProductControllerImpl) Update(c *gin.Context) {
	p.productSvc.Update(c)
}

// Delete godoc
// @Summary Delete product
// @Description Delete product
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} dao.Product
// @Router /products/{id} [delete]
func (p *ProductControllerImpl) Delete(c *gin.Context) {
	p.productSvc.Delete(c)
}

func ProductControllerInit(productSvc service.ProductService) *ProductControllerImpl {
	return &ProductControllerImpl{productSvc: productSvc}
}
