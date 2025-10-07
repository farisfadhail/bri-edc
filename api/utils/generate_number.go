package utils

import (
	"bri-edc/api/config"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

func GenerateTransactionNumber() string {
	timestamp := time.Now().Format("060102150405")
	randomPart := rand.Intn(1000)
	return fmt.Sprintf("TX%s%03d", timestamp, randomPart)
}

func GenerateBatchNumber() string {
	timestamp := time.Now().Format("20060102")

	return fmt.Sprintf("BATCH%s", timestamp)
}

func GenerateHMAC(message string) string {
	secretHmacKey := config.GetEnv("HMAC_SECRET_KEY", "")
	mac := hmac.New(sha256.New, []byte(secretHmacKey))
	mac.Write([]byte(message))
	return hex.EncodeToString(mac.Sum(nil))
}
