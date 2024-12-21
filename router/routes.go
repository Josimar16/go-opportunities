package router

import (
	docs "github.com/Josimar16/go-opportunities/docs"
	"github.com/Josimar16/go-opportunities/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		v1.GET("/opportunities", handler.ListOpportunitiesHandler)
		v1.GET("/opportunities/:id", handler.ShowOpportunitiesHandler)
		v1.POST("/opportunities", handler.CreateOpportunitiesHandler)
		v1.DELETE("/opportunities/:id", handler.RemoveOpportunitiesHandler)
		v1.PUT("/opportunities/:id", handler.EditOpportunitiesHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
