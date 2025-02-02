package controllers

import (
	"net/http"
	"sync"

	"promo_codes_management/config"
	"promo_codes_management/models"
	"promo_codes_management/utils"

	"github.com/gin-gonic/gin"
)

var mu sync.Mutex

// Generate eVouchers API
func GenerateEVouchers(c *gin.Context) {
	var request struct {
		PhoneNumber string `json:"phone_number" binding:"required"`
		Quantity    int    `json:"quantity" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	var promoCodes []models.EVoucher
	for i := 0; i < request.Quantity; i++ {
		promoCode := utils.GenerateUniquePromoCode()
		qrCodeURL := utils.GenerateQRCode(promoCode, promoCode)

		eVoucher := models.EVoucher{
			PromoCode:   promoCode,
			PhoneNumber: request.PhoneNumber,
			QRCodeURL:   qrCodeURL,
		}
		promoCodes = append(promoCodes, eVoucher)
	}

	config.DB.Create(&promoCodes)

	c.JSON(http.StatusOK, gin.H{"message": "eVouchers generated successfully"})
}

// Check Promo Codes API
func CheckPromoCodes(c *gin.Context) {
	var promoCodes []models.EVoucher
	config.DB.Find(&promoCodes)

	c.JSON(http.StatusOK, promoCodes)
}
