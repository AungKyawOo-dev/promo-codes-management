package utils

import (
	"crypto/rand"
	"math/big"
)

// GenerateUniquePromoCode creates an 11-character alphanumeric promo code
func GenerateUniquePromoCode() string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 11)
	for i := range b {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		b[i] = letters[n.Int64()]
	}
	return string(b)
}
