package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.JSON(code, gin.H{
		"message":     msg,
		"status code": code,
	})
}

func sendSuccess(ctx *gin.Context, code int, operation string, data interface{}) {
	ctx.JSON(code, gin.H{
		"status code": code,
		"message":     fmt.Sprintf("operation: %v, successfull!", operation),
		"data":        data,
	})
}
