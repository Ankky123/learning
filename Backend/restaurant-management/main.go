package main

import (
	"os"
	"restaurant-management/middleware"
	"restaurant-management/routes"

	"github.com/gin-gonic/gin"
)

// var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()      //p-func
	router.Use(gin.Logger()) //p-func
	routes.UserRoutes(router)
	router.Use(middleware.Authentication()) //p-func -> Use()

	//initializing the routes
	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)
}
