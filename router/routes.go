package router

import (
	"github.com/gin-gonic/gin"
	"github.com/steixeira93/gopportunities/handler"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/openings/:id", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening/:id", handler.DeleteOpeningHandler)
		v1.PUT("/opening/:id", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
	}
}