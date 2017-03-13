package email

import (
	"net/mail"
	"net/smtp"
	"crypto/tls"
	"fmt"
)

type Sender struct {
	From mail.Address
	Host string
	Port string
	Username string
	Password string
}

func (sender* Sender) SendEmail(message* Message) error {
	content := sender.prepareContent(message)

	auth := smtp.PlainAuth("", sender.Username, sender.Password, sender.Host)

	hostPort := fmt.Sprintf("%s:%s", sender.Host, sender.Port)

	tlsConfig := &tls.Config {
		InsecureSkipVerify: true,
		ServerName: sender.Host,
	}

	conn, err := tls.Dial("tcp", hostPort, tlsConfig)
	if err != nil {
		return err
	}

	c, err := smtp.NewClient(conn, sender.Host)
	if err != nil {
		return err
	}

	if err = c.Auth(auth); err != nil {
		return err
	}

	if err = c.Mail(sender.From.Address); err != nil {
		return err
	}

	if err = c.Rcpt(message.To.Address); err != nil {
		return err
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	if _, err = w.Write([]byte(content)); err != nil {
		return err
	}

	if err = w.Close(); err != nil {
		return err
	}

	c.Quit()

	return nil
}

func (sender* Sender) prepareContent(message* Message) string {
	headers := make(map[string]string)
	headers["From"] = sender.From.String()
	headers["To"] = message.To.String()
	headers["Subject"] = message.Subject

	content := ""
	for k,v := range headers {
		content += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	content += "\r\n" + message.Body

	return content
}