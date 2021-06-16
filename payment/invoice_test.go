package payment

import (
	"context"
	"github.com/Paramosch/predictstock-backend-eng/database"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"os"
	"reflect"
	"strconv"
	"testing"
)

func TestNewInvoice(t *testing.T) {
	ctx := database.Boot(context.Background())
	invoiceService := NewInvoice(ctx)
	if reflect.TypeOf(*invoiceService).Name() != "Invoice" {
		t.Error(reflect.TypeOf(*invoiceService).Name())
		t.Error("Invoice service not instance")
	}
}

func TestCreateInvoice(t *testing.T) {
	ctx := database.Boot(context.Background())

	invoice := NewInvoice(ctx)
	price, _ := strconv.ParseInt(os.Getenv("INVOICE_PRICE"), 0, 64)
	userRepo := database.NewUserRepo(ctx.Value("db").(*gorm.DB))
	user := userRepo.Create(&database.User{TelegramId: 1312312, IsActive: true, IsPaid: true})
	inv, err := invoice.Create(user.GetId(), 10, 10)
	if inv.GetPrice() != price {
		t.Error(err)
		return
	}
	t.Log("Invoice", inv.GetDescription())
}

func TestMarkAsPaid(t *testing.T) {
	testStatus := "paid"

	ctx := database.Boot(context.Background())
	invoiceService := NewInvoice(ctx)
	userRepo := database.NewUserRepo(ctx.Value("db").(*gorm.DB))
	user := userRepo.Create(&database.User{TelegramId: 1312312, IsActive: true, IsPaid: true})
	invoice, _ := invoiceService.Create(user.GetId(), 10, 10)

	inv := invoiceService.MarkAsPaid(invoice)

	if inv.GetStatus() != testStatus {
		t.Errorf("Wrong Status. Expected: %s, Current: %s ", testStatus, inv.GetStatus())
		return
	}

}
