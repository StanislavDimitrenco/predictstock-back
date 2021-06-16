package controllers

import (
	"context"
	"fmt"
	"github.com/Paramosch/predictstock-backend-eng/database"
	"github.com/Paramosch/predictstock-backend-eng/logger"
	"github.com/Paramosch/predictstock-backend-eng/payment"
	"github.com/Paramosch/predictstock-backend-eng/telegram"
	"github.com/Paramosch/predictstock-backend-eng/telegram/messages"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strings"
)

func ResultInvoice(c *fiber.Ctx) error {
	ctx := c.Locals("ctx").(context.Context)
	invoiceService := payment.NewInvoice(ctx)
	fileLog := logger.NewLogger("invoice_requests")
	fileLog.SetFields(logger.Fields{"ip": c.IP()}).LogInfo("New invoice request")

	request := new(payment.RobokassaResult)
	if err := c.BodyParser(request); err != nil {
		fileLog.SetFields(logger.Fields{"ip": c.IP()}).LogError(err)
		return c.Status(412).SendString("Invalid invoice")
	}
	fileLog.SetFields(logger.Fields{"ip": c.IP()}).LogInfo(fmt.Sprintf("Invoice body InvId:%s, OutSum: %s, SignatureValue:%s", request.InvId, request.OutSum, request.SignatureValue))

	if !invoiceService.CheckResult(request) {
		fileLog.SetFields(logger.Fields{"ip": c.IP()}).LogError(fmt.Sprintf("Invalid CheckResult InvId:%s, OutSum: %s, SignatureValue:%s", request.InvId, request.OutSum, request.SignatureValue))
		return c.SendStatus(412)
	}

	invoiceId, price, err := invoiceService.ResultInvoice(request)
	if err != nil {
		fileLog.SetFields(logger.Fields{"ip": c.IP()}).LogError(fmt.Sprintf("Invalid ResultInvoice InvId:%s, OutSum: %s, SignatureValue:%s", request.InvId, request.OutSum, request.SignatureValue))
		return c.Status(412).SendString("Invalid invoice")
	}

	db := ctx.Value("db").(*gorm.DB)
	invoiceRepo := database.NewInvoiceRepo(db)
	invoice, notFound := invoiceRepo.FindBy(map[string]interface{}{"id": invoiceId, "price": price})

	if notFound {
		fileLog.SetFields(logger.Fields{"ip": c.IP()}).LogError(fmt.Sprintf("Invoice notFound InvId:%s, OutSum: %s, SignatureValue:%s", request.InvId, request.OutSum, request.SignatureValue))
		return c.Status(412).SendString("Invalid invoice")
	} else {
		user := invoice.GetUser()
		months := invoice.GetMonths()
		telegramBot := ctx.Value("telegramBot").(*telegram.Bot)
		telegram.NewInvoiceCleaner(ctx, telegramBot).Clean(&user)

		invoice = invoiceService.MarkAsPaid(invoice)
		userRepo := database.NewUserRepo(db)
		userRepo.MarkAsPaid(&user, months)

		fileLog.SetFields(logger.Fields{"invoice": invoice.GetId(), "ip": c.IP()}).LogInfo("Invoice paid")

		expiredAt := strings.ReplaceAll(user.GetPaidUntil().Format("2006-01-02"), "-", "\\-")
		telegram.PushMessage(messages.InvoicePaidSuccess(invoice.GetId(), expiredAt), user.GetTelegramId())
	}

	return c.Status(201).SendString(fmt.Sprintf("OK%d", invoice.GetId()))
}
