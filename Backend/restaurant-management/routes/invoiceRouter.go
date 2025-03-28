package routes

import (
	controller "restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomintRoutes *gin.Engine) {
	incomintRoutes.GET("/invoices", controller.GetInvoices())
	incomintRoutes.GET("/invoices/:invoide_id", controller.GetInvoice())
	incomintRoutes.POST("/invoices", controller.CreateInvoice())
	incomintRoutes.PATCH("/invoices/:invoice_id", controller.UpdateInvoice())
}
