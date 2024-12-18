package handler

import (
	"fmt"
	"github.com/Josimar16/go-opportunities/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RemoveOpportunitiesHandler(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		sendError(c, http.StatusBadRequest, errParamRequired("id", "param").Error())
		return
	}

	opportunity := schemas.Opportunity{}

	if err := db.First(&opportunity, id).Error; err != nil {
		sendError(c, http.StatusNotFound, fmt.Sprintf("opportunity with id: %s not found", id))
		return
	}

	if err := db.Delete(&opportunity).Error; err != nil {
		sendError(c, http.StatusInternalServerError, fmt.Sprintf("error deleting opportunity with id: %s", id))
		return
	}

	sendSuccess(c, http.StatusOK, "remove-opportunity", opportunity)
}
