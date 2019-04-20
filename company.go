package uculqi

import "io/ioutil"

type CompanyInfo struct {
	Name         string
	PhoneSupport string
	EmailSupport string
}

type Company struct {
	Info             *CompanyInfo
	Store            Store
	Charger          Charger
	Messaging        Messaging
	TemplateFilename string
	SendMaxAttempt   int
}

func (c *Company) generateNewInvoice(token string, email string, products []*Product, currency *Currency) (*Invoice, error) {
	invoice, err := newInvoiceWithMinimal(c.Info, &MinimalInformation{
		Email:    email,
		Currency: currency, // by default
		Products: products,
		Token:    token,
	})

	if err != nil {
		return nil, err
	}
	return c.Store.SaveNewInvoice(invoice)
}

func (c *Company) generateNewInvoiceWithOrder(order *Order) (*Invoice, error) {
	invoice, err := newFullInvoice(c.Info, order)
	if err != nil {
		return nil, err
	}
	return c.Store.SaveNewInvoice(invoice)
}

func (c *Company) chargeInvoiceByID(id string) (*Invoice, error) {
	invoice, err := c.Store.GetInvoice(id)
	if err != nil {
		return nil, err
	}

	if invoice.Charged {
		return nil, ErrInvoiceAlreadyCharged
	}

	response, err := c.Charger.ExecuteCharge(invoice)
	if err != nil {
		return nil, err
	}
	invoice.Charged = true // to prevent multiple charges
	invoice, err = c.Store.CheckAsChargedInvoice(invoice)
	if err != nil {
		return nil, err
	}

	invoice, err = c.Store.SetExternalIDOfInvoice(invoice, response.ID)
	if err != nil {
		return nil, err
	}

	return invoice, nil
}

func (c *Company) sendInvoiceAsEmail(id string) (*Invoice, error) {
	invoice, err := c.Store.GetInvoice(id)
	if err != nil {
		return nil, err
	}
	templateData, err := ioutil.ReadFile(c.TemplateFilename)
	if err != nil {
		return nil, err
	}

	for i := 0; i == 0 || i < c.SendMaxAttempt && err != nil; i++ {
		err = c.Messaging.SendInvoiceByEmail(string(templateData), invoice)
	}

	return invoice, err
}
