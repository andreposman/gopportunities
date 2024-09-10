package handler

import (
	"net/http"

	"github.com/andreposman/gopportunities/internal/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}
	opening := schemas.Opening{}

	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.Errf("validation error in update handler: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())

		return
	}

	ID := ctx.Query("id")
	if ID == "" {
		logger.Err("id is empty in update handler")
		sendError(ctx, http.StatusBadRequest, "ID is empty")
	}

	if err := db.First(&opening, ID).Error; err != nil {
		logger.Errf("id not found: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, "opening not found")

		return
	}

	updatedOpening := updateFields(&opening, &request)

	if err := db.Save(&updatedOpening).Error; err != nil {
		logger.Errf("error saving updated opening on db: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, "error saving updated opening")
	}

	sendSuccess(ctx, http.StatusOK, "update-opening", updatedOpening)
}

func updateFields(opening *schemas.Opening, request *UpdateOpeningRequest) *schemas.Opening {
	if request.Role != "" {
		opening.Role = request.Role
	}

	if request.Description != "" {
		opening.Description = request.Description
	}

	if request.Company != "" {
		opening.Company = request.Company
	}

	if request.Location != "" {
		opening.Location = request.Location
	}

	if request.Link != "" {
		opening.Link = request.Link
	}

	if request.IsRemote != nil {
		opening.IsRemote = *request.IsRemote
	}

	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	return opening
}
