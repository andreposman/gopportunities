package handler

import (
	"fmt"
	"net/http"

	"github.com/andreposman/gopportunities/internal/schemas"
	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(ctx *gin.Context) {
	opening := schemas.Opening{}
	ID := ctx.Query("id")
	if ID == "" {
		logger.Err("error finding opening, id is empty")
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("ID", "queryParameter").Error())

		return
	}

	//? get first by id and bind to the &opening for show later
	if err := db.First(&opening, ID).Error; err != nil {
		errMsg := fmt.Sprintf("opening with ID: %s, was not found", ID)

		logger.Errf(errMsg)
		sendError(ctx, http.StatusNotFound, errMsg)

		return
	}

	sendSuccess(ctx, http.StatusOK, "show-opening", opening)
}
