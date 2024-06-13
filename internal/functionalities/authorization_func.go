package functionalities

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"server-bot/internal/models"
	"time"
)

func GenerateTokenForUser(user *models.User) (string, time.Time) {
	// Generate a random 20-byte token
	token := make([]byte, 20)
	_, err := rand.Read(token)
	if err != nil {
		log.Fatal(err)
	}

	// Set token validity for 24 hours
	tokenValidUntil := time.Now().Add(24 * time.Hour)

	// Return the token as a hex string and the validity time
	return hex.EncodeToString(token), tokenValidUntil
}
