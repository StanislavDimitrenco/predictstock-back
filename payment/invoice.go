package payment

import (
	"context"
	"github.com/Paramosch/predictstock-backend-eng/database"
	"gorm.io/gorm"
	"os"
	"strconv"
	"time"
)

type Invoice struct {
	robokassa   *Robokassa
	repo        *database.InvoiceRepo
	price       int64
	description string
}

func NewInvoice(ctx context.Context) *Invoice {
	robokassa := NewRobokassa(
		os.Getenv("ROBOKASSA_ID"),
		os.Getenv("ROBOKASSA_PASSWORD1"),
		os.Getenv("ROBOKASSA_PASSWORD2"),
	)
	db := ctx.Value("db").(*gorm.DB)
	price, _ := strconv.ParseInt(os.Getenv("INVOICE_PRICE"), 0, 64)
	description := os.Getenv("INVOICE_DESCRIPTION")

	return &Invoice{robokassa: robokassa, repo: database.NewInvoiceRepo(db), price: price, description: description}
}

func (i *Invoice) Create(userId int64, mounts int, price int64) (*database.Invoice, string) {
	invoice := &database.Invoice{
		UserID:      userId,
		Status:      "unpaid",
		Price:       price,
		Months:      mounts,
		Description: i.description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		PaidAt:      time.Now(),
	}
	i.repo.Create(invoice)
	isTest, _ := strconv.ParseBool(os.Getenv("IS_TEST"))
	url := i.robokassa.URL(invoice.GetId(), invoice.GetPrice(), invoice.GetDescription(), isTest)
	return invoice, url
}

func (i *Invoice) GetOldInvoices(userId int64) []database.Invoice {
	invoices := i.repo.Filter(map[string]interface{}{"user_id": userId, "status": "unpaid", "deleted_at": nil})
	return invoices
}

func (i *Invoice) RemoveInvoice(invoice *database.Invoice) {
	i.repo.Delete(invoice)
}

func (i *Invoice) AttachMessageId(invoice *database.Invoice, id int) *database.Invoice {
	invoice.MessageId = id
	return i.repo.Save(invoice)
}

func (i *Invoice) CheckResult(r *RobokassaResult) bool {
	return i.robokassa.CheckResult(r)
}

func (i *Invoice) ResultInvoice(r *RobokassaResult) (int, int, error) {
	return i.robokassa.ResultInvoice(r)
}

func (i *Invoice) FindById(invoiceId int64) (database.Invoice, bool) {
	return i.repo.FindById(invoiceId)
}
func (i *Invoice) MarkAsPaid(invoice *database.Invoice) *database.Invoice {
	invoice.Status = "paid"
	invoice.PaidAt = time.Now()
	i.repo.Save(invoice)

	return invoice
}
