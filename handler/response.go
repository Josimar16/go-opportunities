package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func sendError(c *gin.Context, code int, message string) {
	c.Header("Content-Type", "application/json")
	c.JSON(code, gin.H{
		"message": message,
		"status":  code,
	})
}

func sendSuccess(c *gin.Context, code int, op string, data interface{}) {
	c.Header("Content-Type", "application/json")
	c.JSON(code, gin.H{
		"message": fmt.Sprintf("operation from handler %s successfull", op),
		"data":    data,
		"status":  code,
	})
}
