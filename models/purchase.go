package models

import (
	"gorm.io/gorm"
)

// Purchase History
type Purchase struct {
	gorm.Model
	UserID     uint
	Amount     float64
	PromoCodes []EVoucher `gorm:"foreignKey:PhoneNumber;references:PhoneNumber"`
}
