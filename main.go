package main

import (
	"vigil/email"
	"net/mail"
	"log"
	"os"
	"fmt"

	"gopkg.in/pg.v5"
)

func main() {
	db := pg.Connect(&pg.Options{
		Addr: "vigil_postgres:5432",
    User:     "postgres",
    Password: "",
    Database: "postgres",
	})

	var n int
	_, err := db.QueryOne(pg.Scan(&n), "SELECT 1")
	if err != nil {
		panic(err)
	}
	fmt.Println(n)

	err = db.Close()
	if err != nil {
		panic(err)
	}

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
