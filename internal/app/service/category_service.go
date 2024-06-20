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

type CategoryService interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type CategoryServiceImpl struct {
	categoryRepo repository.CategoryRepository
}

func (catSvc CategoryServiceImpl) GetAll(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	data, err := catSvc.categoryRepo.GetAll(limit, offset)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.InvalidRequest, err, pkg.Null()))
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (catSvc CategoryServiceImpl) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := catSvc.categoryRepo.GetByID(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, pkg.BuildResponse(constant.DataNotFound, err, pkg.Null()))
		}
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (catSvc CategoryServiceImpl) Create(c *gin.Context) {
	var category dao.Category
	_ = c.BindJSON(&category)
	data, err := catSvc.categoryRepo.Create(category)
	if err != nil {
		if errors.Is(err, gorm.ErrInvalidData) {
			c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.InvalidRequest, err, pkg.Null()))
			return
		}
	}

	c.JSON(http.StatusCreated, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (catSvc CategoryServiceImpl) Update(c *gin.Context) {
	var category dao.Category
	_ = c.BindJSON(&category)
	data, err := catSvc.categoryRepo.Update(category)
	if err != nil {
		if errors.Is(err, gorm.ErrInvalidData) {
			c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.InvalidRequest, err, pkg.Null()))
			return
		}
	}

	c.JSON(http.StatusCreated, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (catSvc CategoryServiceImpl) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := catSvc.categoryRepo.Delete(id)
	if err != nil {
		if errors.Is(err, gorm.ErrInvalidData) {
			c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.InvalidRequest, err, pkg.Null()))
			return
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, pkg.BuildResponse(constant.DataNotFound, err, pkg.Null()))
			return
		}
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), pkg.Null()))
}

func CategoryServiceInit(cateRepo repository.CategoryRepository) *CategoryServiceImpl {
	return &CategoryServiceImpl{
		categoryRepo: cateRepo,
	}
}
