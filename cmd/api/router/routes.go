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
