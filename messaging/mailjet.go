package messaging

import (
	"gopkg.in/gomail.v2"
)

type MailJetBackend struct {
	d         *gomail.Dialer
	server    string
	port      int
	email     string
	password  string
	fromName  string
	fromEmail string
	maxTries  int
}

func NewMailJetBackend(mailJetEmail, mailJetPassword, fromName, fromEmail string) *MailJetBackend {
	backend := &MailJetBackend{
		server:    "in-v3.mailjet.com",
		port:      587,
		email:     mailJetEmail,
		password:  mailJetPassword,
		fromName:  fromName,
		fromEmail: fromEmail,
		maxTries:  4,
	}
	d := gomail.NewDialer(backend.server, backend.port, backend.email, backend.password)
	backend.d = d
	return backend
}
