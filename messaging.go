package uculqi

type Messaging interface {
	SendInvoiceByEmail(templateData string, invoice *Invoice) error
}
