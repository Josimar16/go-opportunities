package handler

import (
	"fmt"
	"github.com/Josimar16/go-opportunities/schemas"
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

type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type CreateOpportunityResponse struct {
	Message string                      `json:"message"`
	Status  int                         `json:"status"`
	Data    schemas.OpportunityResponse `json:"data"`
}

type RemoveOpportunityResponse struct {
	Message string                      `json:"message"`
	Status  int                         `json:"status"`
	Data    schemas.OpportunityResponse `json:"data"`
}

type ListOpportunityResponse struct {
	Message string                        `json:"message"`
	Status  int                           `json:"status"`
	Data    []schemas.OpportunityResponse `json:"data"`
}

type ShowOpportunityResponse struct {
	Message string                      `json:"message"`
	Status  int                         `json:"status"`
	Data    schemas.OpportunityResponse `json:"data"`
}

type EditOpportunityResponse struct {
	Message string                      `json:"message"`
	Status  int                         `json:"status"`
	Data    schemas.OpportunityResponse `json:"data"`
}
