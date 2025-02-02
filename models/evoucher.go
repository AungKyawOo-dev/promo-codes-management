package models

import (
	"gorm.io/gorm"
)

// EVoucher Model
type EVoucher struct {
	gorm.Model
	PromoCode   string `gorm:"uniqueIndex"`
	PhoneNumber string
	QRCodeURL   string
}
