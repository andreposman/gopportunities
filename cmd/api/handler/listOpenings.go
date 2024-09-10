package handler

import (
	"net/http"

	"github.com/andreposman/gopportunities/internal/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listenign openings")
		return
	}

	sendSuccess(ctx, http.StatusOK, "list-openings", openings)

}
