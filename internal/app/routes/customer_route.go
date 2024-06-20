package routes

import (
	"awesomeProject/config"
	"awesomeProject/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func CustomerRoute(init *config.Initialize, g *gin.RouterGroup) *gin.RouterGroup {
	g.GET("", middleware.AuthMiddleware(), init.CustomerCtrl.GetAll)
	g.POST("", init.CustomerCtrl.Create)
	return g
}
