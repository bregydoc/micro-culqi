package uculqi

import "errors"

var ErrInvalidInvoice = errors.New("invalid invoice, please check your information")
var ErrInvoiceNotCharged = errors.New("your invoice are not charged yet")
var ErrInvoiceAlreadyCharged = errors.New("your invoice already was charged")
var ErrInvoiceNotFound = errors.New("invoice not found")
