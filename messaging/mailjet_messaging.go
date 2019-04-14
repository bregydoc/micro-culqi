package messaging

import (
	"bytes"
	"fmt"
	"github.com/bregydoc/micro-culqi"
	"gopkg.in/gomail.v2"
	"text/template"
)

func (mail *MailJetBackend) sendMail(data string, params interface{}, subject string, to []string) error {
	from := mail.fromName + " <" + mail.fromEmail + ">"
	body := new(bytes.Buffer)
	temp, err := template.New("invoice").Parse(data)
	if err != nil {
		return err
	}
	err = temp.Execute(body, params)
	if err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body.String())
	// m.Attach("/home/Alex/cat.jpg") // ready to attachment

	if err := mail.d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func (mail *MailJetBackend) SendInvoiceByEmail(templateData string, invoice *uculqi.Invoice) error {
	if !invoice.IsValid() {
		return uculqi.ErrInvalidInvoice
	}

	if !invoice.Charged {
		return uculqi.ErrInvoiceNotCharged
	}

	subject := fmt.Sprintf("Your invoice %s", invoice.ID)
	return mail.sendMail(templateData, invoice.ToFlatten(), subject, []string{invoice.Email})
}
