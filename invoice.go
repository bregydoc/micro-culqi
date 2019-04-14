package uculqi

import (
	"regexp"
	"time"
)

type Invoice struct {
	ID         string
	Order      *Order
	Email      string
	Charged    bool
	Company    *CompanyInfo
	ChargedAt  time.Time
	ExternalID string
}

func (invoice *Invoice) IsValid() bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if len(invoice.ID) < 5 {
		return false
	}

	if invoice.Order == nil {
		return false
	}

	if re.MatchString(invoice.Email) {
		return false
	}

	if len(invoice.Order.ID) < 5 {
		return false
	}

	if invoice.Order.Currency == nil {
		return false
	}

	if invoice.Order.Info == nil {
		return false
	}

	if invoice.Order.Token == "" {
		return false
	}

	if len(invoice.Order.Products) == 0 {
		return false
	}

	if invoice.Order.Info.Email == "" {
		return false
	}

	return true
}
