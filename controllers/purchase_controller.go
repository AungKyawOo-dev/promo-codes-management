package controllers

import (
	"net/http"

	"promo_codes_management/config"
	"promo_codes_management/models"

	"github.com/gin-gonic/gin"
)

// Process Payment API
func ProcessPayment(c *gin.Context) {
	var request struct {
		UserID uint    `json:"user_id" binding:"required"`
		Amount float64 `json:"amount" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var availableCodes []models.EVoucher
	config.DB.Limit(int(request.Amount / 10)).Find(&availableCodes)

	if len(availableCodes) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No promo codes available"})
		return
	}

	// Update promo codes to be associated with the user
	for _, code := range availableCodes {
		config.DB.Model(&code).Update("PhoneNumber", request.UserID)
	}

	// Store purchase record
	purchase := models.Purchase{
		UserID:     request.UserID,
		Amount:     request.Amount,
		PromoCodes: availableCodes,
	}
	config.DB.Create(&purchase)

	c.JSON(http.StatusOK, gin.H{"message": "Payment processed and promo codes assigned"})
}
