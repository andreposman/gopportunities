package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", pingHandler)
		v1.GET("/opening", openingHandler)
		v1.POST("/opening", createOpeningHandler)
		v1.PUT("/opening", updateOpeningHandler)
		v1.DELETE("/opening", deleteOpeningHandler)
	}
}

func pingHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func openingHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET Opening",
	})
}

func createOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "POST Opening",
	})
}

func updateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "PUT Opening",
	})
}

func deleteOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "DELETE Opening",
	})
}
