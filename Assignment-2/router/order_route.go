package router

import (
	"github.com/gin-gonic/gin"
	"github.com/roihan12/assignment-2/controller"
	"github.com/roihan12/assignment-2/repository"
	"github.com/roihan12/assignment-2/service"
	"gorm.io/gorm"
)

func SetupOrderRoute(router *gin.Engine, db *gorm.DB) {
	orderRepository := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepository)
	orderController := controller.NewOrderController(orderService)

	orderRouter := router.Group("/api/orders")
	{
		orderRouter.POST("/", orderController.CreateOrder)
		orderRouter.GET("/:id", orderController.GetOrderById)
		orderRouter.PUT("/:id", orderController.UpdateOrder)
		orderRouter.GET("/", orderController.GetAllOrder)
		orderRouter.DELETE("/:id", orderController.DeleteOrder)

	}
}
