package routes

import (
	"awesomeProject/config"
	_ "awesomeProject/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"
)

func Route(init *config.Initialize) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		MaxAge:       12 * time.Hour,
	}))

	api := router.Group("/api/v1")
	{
		CustomerRoute(init, api.Group("/customers"))
		AuthRoute(init, api.Group("/auth"))
		CategoryRoute(init, api.Group("/categories"))
		ProductRoute(init, api.Group("/products"))
	}

	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
