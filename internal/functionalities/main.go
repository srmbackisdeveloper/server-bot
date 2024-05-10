package functionalities

import (
	"encoding/json"
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

func WriteJSON(w http.ResponseWriter, status int, anything interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(anything)
	if err != nil {
		return
	}
}

func GenerateCode() string {
	code := fmt.Sprintf("%04d", rand.Intn(10000))
	return code
}

func SendVerificationCodeEmail(to, code string) {
	// Setup
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	smtpUser := os.Getenv("FROM_EMAIL")
	smtpPassword := os.Getenv("FROM_EMAIL_PASSWORD")

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPassword)

	// Email
	m := gomail.NewMessage()
	m.SetHeader("From", smtpUser)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Necompany - Verification Code")
	m.SetBody("text/plain", "Here is your verification code: "+code)

	// Sending...
	if err := d.DialAndSend(m); err != nil {
		log.Printf("Error sending email: %v\n", err)
	}
}
