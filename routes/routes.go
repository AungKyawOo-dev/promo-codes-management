package routes

import (
	"promo_codes_management/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	// eVoucher Routes
	router.POST("/generate-evoucher", controllers.GenerateEVouchers)
	router.GET("/check-promocodes", controllers.CheckPromoCodes)

	// Payment & Purchase
	router.POST("/process-payment", controllers.ProcessPayment)

	return router
}
