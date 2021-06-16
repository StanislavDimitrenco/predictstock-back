package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Paramosch/predictstock-backend-eng/database"
	"github.com/Paramosch/predictstock-backend-eng/payment"
	"github.com/Paramosch/predictstock-backend-eng/providers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestResultInvoice(t *testing.T) {
	ctx := providers.Boot(context.Background())
	app := ctx.Value("fiber").(*fiber.App)

	invoiceService := payment.NewInvoice(ctx)
	userRepo := database.NewUserRepo(ctx.Value("db").(*gorm.DB))
	user := userRepo.Create(&database.User{TelegramId: 1312312, IsActive: true, IsPaid: true})
	invoice, _ := invoiceService.Create(user.GetId())
	fmt.Println(invoice)

	signatureValue := "98A81B9E2E6CE063901680A93D9AD8B5"

	app.Post("/robokassa/result", ResultInvoice)

	dataMap, _ := json.Marshal(map[string]interface{}{
		"OutSum":         "100.000000",
		"InvId":          "76",
		"SignatureValue": signatureValue,
	})

	req := httptest.NewRequest("POST", "/robokassa/result", bytes.NewBuffer(dataMap))
	req.Header.Set("Content-Type", "application/json")

	// http.Response
	resp, _ := app.Test(req)

	if resp == nil {
		t.Error("Undefined response")
	} else if resp.StatusCode == 201 {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	} else {
		t.Errorf("Wrond Status Code: %d, expected: 200", resp.StatusCode)
	}
}
