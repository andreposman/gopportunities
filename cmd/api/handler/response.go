package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// ? this is to ensure that the response follow the order that I want
type CustomSuccessResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.JSON(code, gin.H{
		"message":     msg,
		"status code": code,
	})
}

func sendSuccess(ctx *gin.Context, code int, operation string, data interface{}) {
	response := CustomSuccessResponse{
		StatusCode: code,
		Message:    fmt.Sprintf("operation: %v, successfull!", operation),
		Data:       data,
	}
	ctx.JSON(response.StatusCode, response)
}
