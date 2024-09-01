package router

import (
	"github.com/andreposman/gopportunities/cmd/api/handler"
	"github.com/gin-gonic/gin"
)

func initRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", handler.PingHandler)

		v1.GET("/opening", handler.OpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)

		v1.GET("/openings", handler.ListOpeningsHandler)
	}
}
