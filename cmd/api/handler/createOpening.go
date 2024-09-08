package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.Errf("error validating creating opening request: %v", err.Error())
		errMsg := err.Error()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errMsg,
		})
		
		return
	}

	if err := db.Create(&request).Error; err != nil {
		logger.Errf("error creating opening: %v", err)
		return
	}

	logger.Infof("Request received: %+v", request)

}
