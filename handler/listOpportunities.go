package handler

import (
	"github.com/Josimar16/go-opportunities/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListOpportunitiesHandler(c *gin.Context) {
	opportunities := []schemas.Opportunity{}

	if err := db.Find(&opportunities).Error; err != nil {
		sendError(c, http.StatusInternalServerError, "Error listing opportunities on database")
		return
	}

	sendSuccess(c, http.StatusOK, "list-opportunities", opportunities)
}
