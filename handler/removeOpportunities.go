package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RemoveOpportunitiesHandler(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "Hello World",
	})
}
