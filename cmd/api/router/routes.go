package router

import (
	"github.com/andreposman/gopportunities/cmd/api/handler"

	docs "github.com/andreposman/gopportunities/docs"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initRoutes(router *gin.Engine) {
	handler.InitHandler()
	basePathV1 := "/api/v1"
	docs.SwaggerInfo.BasePath = basePathV1

	v1 := router.Group(basePathV1)
	{
		v1.GET("/ping", handler.PingHandler)

		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)

		v1.GET("/openings", handler.ListOpeningsHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
