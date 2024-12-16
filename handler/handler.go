package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOpportunitiesHandler(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "Hello World",
	})
}

func ShowOpportunitiesHandler(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "Hello World",
	})
}

func ListOpportunitiesHandler(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "Hello World",
	})
}

func RemoveOpportunitiesHandler(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "Hello World",
	})
}

func EditOpportunitiesHandler(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "Hello World",
	})
}
