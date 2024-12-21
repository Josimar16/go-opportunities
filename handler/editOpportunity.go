package handler

import (
	"fmt"
	"github.com/Josimar16/go-opportunities/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func EditOpportunitiesHandler(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		sendError(c, http.StatusBadRequest, errParamRequired("id", "param").Error())
		return
	}

	request := UpdateOpportunityRequest{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Error validating request: %v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	opportunity := schemas.Opportunity{}

	if err := db.First(&opportunity, id).Error; err != nil {
		sendError(c, http.StatusNotFound, fmt.Sprintf("opportunity with id: %s not found", id))
		return
	}

	if request.Role != "" {
		opportunity.Role = request.Role
	}

	if request.Company != "" {
		opportunity.Company = request.Company
	}

	if request.Location != "" {
		opportunity.Location = request.Location
	}

	if request.Remote != nil {
		opportunity.Remote = *request.Remote
	}

	if request.Link != "" {
		opportunity.Link = request.Link
	}

	if request.Salary > 0 {
		opportunity.Salary = request.Salary
	}

	if err := db.Save(&opportunity).Error; err != nil {
		sendError(c, http.StatusInternalServerError, fmt.Sprintf("error updating opportunity with id: %s", id))
		return
	}

	sendSuccess(c, http.StatusOK, "edit-opportunity", opportunity)
}
