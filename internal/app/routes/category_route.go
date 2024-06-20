package routes

import (
	"awesomeProject/config"
	"github.com/gin-gonic/gin"
)

func CategoryRoute(init *config.Initialize, g *gin.RouterGroup) *gin.RouterGroup {
	g.GET("", init.CategoryCtrl.GetAll)
	g.POST("", init.CategoryCtrl.Create)
	g.DELETE("/:id", init.CategoryCtrl.Delete)
	return g
}
