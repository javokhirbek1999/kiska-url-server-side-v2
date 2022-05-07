package utils

import (
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func SendMail(toEmail, message string) error {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	from := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")

	to := []string{toEmail}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	newMessage := []byte(message)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	if err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, newMessage); err != nil {
		return err
	}

	return nil
}
