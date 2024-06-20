package routes

import (
	"awesomeProject/config"
	"github.com/gin-gonic/gin"
)

func ProductRoute(init *config.Initialize, g *gin.RouterGroup) *gin.RouterGroup {
	g.GET("", init.ProductCtrl.GetAll)
	g.GET("/:id", init.ProductCtrl.GetByID)
	g.POST("", init.ProductCtrl.Create)
	g.PUT("/:id", init.ProductCtrl.Update)
	g.DELETE("/:id", init.ProductCtrl.Delete)
	return g
}
