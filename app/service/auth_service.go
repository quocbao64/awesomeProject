package service

import (
	"awesomeProject/app/auth"
	"awesomeProject/app/constant"
	"awesomeProject/app/pkg"
	"awesomeProject/app/repository"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type AuthService interface {
	Login(c *gin.Context)
}

type AuthServiceImpl struct {
	customerRepo repository.CustomerRepository
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (authSvc AuthServiceImpl) Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	data, err := authSvc.customerRepo.GetByEmail(email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, pkg.BuildResponse(constant.DataNotFound, "Email not found", pkg.Null()))
			return
		}
	}

	if pkg.ComparePassword(password, data.Password) {
		tokenString, err := auth.GenerateJWT(email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, pkg.BuildResponse(constant.Success, "Error creating token", data))
			return
		}
		response := LoginResponse{AccessToken: tokenString, RefreshToken: "dfadsaf"}
		c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), response))
	} else {
		c.JSON(http.StatusUnauthorized, pkg.BuildResponse(constant.Unauthorized, "Unauthorized", pkg.Null()))
	}
}

func AuthServiceInit(customerRepo repository.CustomerRepository) *AuthServiceImpl {
	return &AuthServiceImpl{customerRepo: customerRepo}
}
