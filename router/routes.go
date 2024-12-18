package router

import (
	"github.com/Josimar16/go-opportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/opportunities", handler.ListOpportunitiesHandler)
		v1.GET("/opportunities/{id}", handler.ShowOpportunitiesHandler)
		v1.POST("/opportunities", handler.CreateOpportunitiesHandler)
		v1.DELETE("/opportunities", handler.RemoveOpportunitiesHandler)
		v1.PUT("/opportunities", handler.EditOpportunitiesHandler)
	}
}
