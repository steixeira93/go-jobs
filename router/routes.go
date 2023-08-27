package router

import (
	"github.com/gin-gonic/gin"
	"github.com/steixeira93/gopportunities/handler"
	docs "github.com/steixeira93/gopportunities/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {

	//Initialize Handler
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	
	v1 := router.Group(basePath)
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}