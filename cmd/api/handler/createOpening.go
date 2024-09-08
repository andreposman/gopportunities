package handler

import (
	"net/http"

	"github.com/andreposman/gopportunities/internal/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {

	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.Errf("error validating creating opening request: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())

		return
	}

	//? converting to sqlite valid data
	opening := schemas.Opening{
		Role:        request.Role,
		Description: request.Description,
		Company:     request.Company,
		Location:    request.Location,
		IsRemote:    *request.IsRemote,
		Link:        request.Link,
		Salary:      request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errf("error creating opening in database: %v", err)
		sendError(ctx, http.StatusInternalServerError, "error saving to the database")
		return
	}

	logger.Infof("Request received: %+v", request)
	sendSuccess(ctx, http.StatusCreated, "createOpening", opening)

}
