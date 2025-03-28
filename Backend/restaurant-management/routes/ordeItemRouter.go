package routes

import (
	controller "restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems/", controller.GetOrderItems())
	incomingRoutes.GET("/orderItems/orderItem_id", controller.GetOrderItem())
	incomingRoutes.GET("orderItems-order/:order_id", controller.GetOrderItemsByOrder())
	incomingRoutes.POST("/orderItem", controller.CreateOrderItem())
	incomingRoutes.PATCH("/orderItem/:orderItem_id", controller.UpdateOrderItem())
}
