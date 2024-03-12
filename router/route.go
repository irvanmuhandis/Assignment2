package router

import (
	"assignment2/controllers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func MyRouter() *gin.Engine {
	//declare router and controller
	router := gin.Default()
	controller := controllers.OrderController{}

	//url route
	order := router.Group("/orders")
	{
		order.GET("", controller.GetOrders)
		order.PUT(":orderId", controller.UpdateOrder)
		order.POST("", controller.CreateOrder)
		order.DELETE(":orderId", controller.DeleteOrder)
	}

	// add swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
