package main

import (
	"vigil/email"
	"net/mail"
	"log"
	"os"
)

func main() {
	message := email.Message{
		To: mail.Address{Address: os.Getenv("VIGIL_TO")},
		Subject: "Golang test",
		Body: "This is test",
	}

	sender := email.Sender{
		From: mail.Address{Name: os.Getenv("VIGIL_FROM_NAME"), Address: os.Getenv("VIGIL_FROM_ADDRESS")},
		Host: os.Getenv("VIGIL_HOST"),
		Port: os.Getenv("VIGIL_PORT"),
		Username: os.Getenv("VIGIL_USERNAME"),
		Password: os.Getenv("VIGIL_PASSWORD"),
	}

	if err := sender.SendEmail(&message); err != nil {
		log.Panic(err)
	}
}
