package handler

import (
	"github.com/Josimar16/go-opportunities/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOpportunitiesHandler(c *gin.Context) {
	request := CreateOpportunityRequest{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Error validating request: %v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	opportunity := schemas.Opportunity{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opportunity).Error; err != nil {
		logger.Errorf("Error creating opportunity: %v", err)
		sendError(c, http.StatusInternalServerError, "Error creating opportunity on database")
		return
	}

	sendSuccess(c, http.StatusCreated, "create-opportunity", opportunity)
}
