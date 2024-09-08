package handler

import (
	"fmt"
	"net/http"

	"github.com/andreposman/gopportunities/internal/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(ctx *gin.Context) {

	opening := schemas.Opening{}
	ID := ctx.Query("id")
	if ID == "" {
		logger.Err("error deleting opening, id is empty")
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("ID", "queryParameter").Error())

		return
	}

	//? get first by id and bind to the &opening for deleting later
	if err := db.First(&opening, ID).Error; err != nil {
		errMsg := fmt.Sprintf("opening with ID: %s, was not found", ID)

		logger.Errf(errMsg)
		sendError(ctx, http.StatusNotFound, errMsg)

		return
	}

	//? deleting binded opening
	if err := db.Delete(&opening).Error; err != nil {
		errMsg := fmt.Sprintf("error deleting the opening with ID: %s", ID)

		logger.Errf(errMsg)
		sendError(ctx, http.StatusInternalServerError, errMsg)
		return
	}

	sendSuccess(ctx, http.StatusOK, "delete opening", opening)
}
