package controller

import (
	"awesomeProject/app/service"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(c *gin.Context)
}

type AuthControllerImpl struct {
	authService service.AuthService
}

// Login godoc
// @Summary Login
// @Description Login
// @Tags auth
// @Accept mpfd
// @Produce json
// @Param email formData string true "email"
// @Param password formData string true "password"
// @Success 200 {object} service.LoginResponse
// @Router /auth/login [post]
func (a AuthControllerImpl) Login(c *gin.Context) {
	a.authService.Login(c)
}

func AuthControllerInit(authService service.AuthService) *AuthControllerImpl {
	return &AuthControllerImpl{authService: authService}
}
