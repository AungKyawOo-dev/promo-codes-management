package utils

import (
	"fmt"
	"log"

	"github.com/skip2/go-qrcode"
)

// GenerateQRCode generates a QR code and saves it to a file
func GenerateQRCode(content, filename string) string {
	filePath := fmt.Sprintf("qrcodes/%s.png", filename)
	err := qrcode.WriteFile(content, qrcode.Medium, 256, filePath)
	if err != nil {
		log.Println("Error generating QR Code:", err)
		return ""
	}
	return filePath
}
