package handler

import (
	"fmt"
	"github.com/Josimar16/go-opportunities/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowOpportunitiesHandler(c *gin.Context) {
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

	sendSuccess(c, http.StatusOK, "show-opportunity", opportunity)
}
