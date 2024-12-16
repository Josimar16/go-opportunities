package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opportunities", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello World",
			})
		})
		v1.GET("/opportunities/{id}", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello World",
			})
		})
		v1.POST("/opportunities", func(c *gin.Context) {
			c.JSON(http.StatusCreated, gin.H{
				"message": "Hello World",
			})
		})
		v1.DELETE("/opportunities", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello World",
			})
		})
		v1.PUT("/opportunities", func(c *gin.Context) {
			c.JSON(http.StatusNoContent, gin.H{
				"message": "Hello World",
			})
		})
	}
}
