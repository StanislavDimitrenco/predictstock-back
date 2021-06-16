package database

import (
	"errors"
	"gorm.io/gorm"
)

type InvoiceRepo struct {
	db *gorm.DB
}

func NewInvoiceRepo(db *gorm.DB) *InvoiceRepo {
	return &InvoiceRepo{db: db.Preload("User")}
}

func (repo InvoiceRepo) Create(invoice *Invoice) *Invoice {
	repo.db.Create(&invoice)
	//TODO: Реализовать это более изящным способом
	repo.db.Find(&invoice)
	return invoice
}

func (repo InvoiceRepo) Save(invoice *Invoice) *Invoice {
	repo.db.Save(&invoice)
	return invoice
}

func (repo InvoiceRepo) Update(id int64, invoice *Invoice) *Invoice {
	repo.db.Where("id=", id).Updates(invoice)
	return invoice
}

func (repo InvoiceRepo) Delete(invoice *Invoice) {
	repo.db.Delete(invoice)
}

func (repo InvoiceRepo) FindById(id int64) (Invoice, bool) {
	var invoice Invoice
	result := repo.db.Preload("User").First(&invoice, id)
	return invoice, errors.Is(result.Error, gorm.ErrRecordNotFound)
}

func (repo InvoiceRepo) FindBy(param map[string]interface{}) (*Invoice, bool) {
	var invoice Invoice
	result := repo.db.Where(param).First(&invoice)
	return &invoice, errors.Is(result.Error, gorm.ErrRecordNotFound)
}

func (repo InvoiceRepo) Filter(param map[string]interface{}) []Invoice {
	var invoice []Invoice
	repo.db.Where(param).Find(&invoice)
	return invoice
}
