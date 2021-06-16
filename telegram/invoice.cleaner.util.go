package telegram

import (
	"context"
	"fmt"
	"github.com/Paramosch/predictstock-backend-eng/database"
	"github.com/Paramosch/predictstock-backend-eng/payment"
	"strconv"
)

type InvoiceCleaner struct {
	ctx context.Context
	bot *Bot
}

func NewInvoiceCleaner(ctx context.Context, bot *Bot) *InvoiceCleaner {
	return &InvoiceCleaner{ctx: ctx, bot: bot}
}

func (i InvoiceCleaner) Clean(userId *database.User) {
	invoiceService := payment.NewInvoice(i.ctx)
	invoices := invoiceService.GetOldInvoices(userId.GetId())
	for _, invoice := range invoices {
		invoiceService.RemoveInvoice(&invoice)
		err := i.bot.Client.DeleteMessage(strconv.Itoa(userId.GetTelegramId()), invoice.GetMessageId())
		if err != nil {
			fmt.Println(err)
		}
	}
}
