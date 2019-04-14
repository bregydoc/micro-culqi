package uculqi

type Store interface {
	SaveNewInvoice(invoice *Invoice) (*Invoice, error)
	GetInvoice(ID string) (*Invoice, error)
	CheckAsChargedInvoice(invoice *Invoice) (*Invoice, error)
	SetExternalIDOfInvoice(invoice *Invoice, externalID string) (*Invoice, error)
}
