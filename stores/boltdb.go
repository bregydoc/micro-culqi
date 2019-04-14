package stores

import (
	"github.com/asdine/storm"
	"github.com/bregydoc/micro-culqi"
	"github.com/bregydoc/micro-culqi/littleid"
	"time"
)

type BoltStore struct {
	filePath string
	db       *storm.DB
}

func NewBoltStore(filepath string) (*BoltStore, error) {
	db, err := storm.Open(filepath)
	if err != nil {
		return nil, err
	}

	return &BoltStore{
		db:       db,
		filePath: filepath,
	}, nil
}

func (store *BoltStore) SaveNewInvoice(invoice *uculqi.Invoice) (*uculqi.Invoice, error) {
	if invoice.ID == "" {
		invoice.ID = littleid.New()
	}

	err := store.db.Save(invoice)
	if err != nil {
		return nil, err
	}

	return invoice, err
}

func (store *BoltStore) GetInvoice(ID string) (*uculqi.Invoice, error) {
	invoice := new(uculqi.Invoice)
	err := store.db.One("ID", ID, invoice)
	if err != nil {
		return nil, uculqi.ErrInvoiceNotFound
	}
	return invoice, nil
}

func (store *BoltStore) CheckAsChargedInvoice(invoice *uculqi.Invoice) (*uculqi.Invoice, error) {
	invoice, err := store.GetInvoice(invoice.ID)
	if err != nil {
		return nil, err
	}

	if err = store.db.UpdateField(invoice, "Charged", true); err != nil {
		return nil, err
	}

	if err = store.db.UpdateField(invoice, "ChargedAt", time.Now()); err != nil {
		return nil, err
	}
	return invoice, err

}

func (store *BoltStore) SetExternalIDOfInvoice(invoice *uculqi.Invoice, externalID string) (*uculqi.Invoice, error) {
	invoice, err := store.GetInvoice(invoice.ID)
	if err != nil {
		return nil, err
	}

	if err = store.db.UpdateField(invoice, "ExternalID", externalID); err != nil {
		return nil, err
	}

	return invoice, err
}
