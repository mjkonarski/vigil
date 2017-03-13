package email

import "net/mail"

type Message struct {
	To mail.Address
	Subject string
	Body string
}