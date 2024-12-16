package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func EditOpportunitiesHandler(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "Hello World",
	})
}
