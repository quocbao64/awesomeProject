package controller

import (
	"awesomeProject/internal/app/service"
	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type CategoryControllerImpl struct {
	categorySvc service.CategoryService
}

// GetAll godoc
// @Summary Get all categories
// @Description Get all categories
// @Tags categories
// @Accept json
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} dao.Category
// @Router /categories [get]
func (cateCtrl CategoryControllerImpl) GetAll(c *gin.Context) {
	cateCtrl.categorySvc.GetAll(c)
}

// GetByID godoc
// @Summary Get category by ID
// @Description Get category by ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} dao.Category
// @Router /categories/{id} [get]
func (cateCtrl CategoryControllerImpl) GetByID(c *gin.Context) {
	cateCtrl.categorySvc.GetByID(c)
}

// Create godoc
// @Summary Create category
// @Description Create category
// @Tags categories
// @Accept json
// @Produce json
// @Param name body string true "Category Name"
// @Success 200 {object} dao.Category
// @Router /categories [post]
func (cateCtrl CategoryControllerImpl) Create(c *gin.Context) {
	cateCtrl.categorySvc.Create(c)
}

// Delete godoc
// @Summary Delete category
// @Description Delete category
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200
// @Router /categories/{id} [delete]
func (cateCtrl CategoryControllerImpl) Delete(c *gin.Context) {
	cateCtrl.categorySvc.Delete(c)
}

// Update godoc
// @Summary Update category
// @Description Update category
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param name body string true "Category Name"
// @Success 200 {object} dao.Category
// @Router /categories/{id} [put]
func (cateCtrl CategoryControllerImpl) Update(c *gin.Context) {
	cateCtrl.categorySvc.Update(c)
}

func CategoryControllerInit(categoryService service.CategoryService) *CategoryControllerImpl {
	return &CategoryControllerImpl{
		categorySvc: categoryService,
	}
}
