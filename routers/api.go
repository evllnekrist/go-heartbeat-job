package routers

import (
	"github.com/gin-gonic/gin"
	"go-heartbeat-job/controllers"
)

func StartServer() *gin.Engine{
	router := gin.Default()

	orderRouter := router.Group("/weather")
	{
		orderRouter.GET("/:orderId", controllers.WeatherGet)
		orderRouter.GET("/", controllers.WeatherGetList)
		orderRouter.POST("/", controllers.WeatherCreate)
	}

	return router
}