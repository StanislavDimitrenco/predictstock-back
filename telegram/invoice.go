package telegram

import (
	"context"
	"github.com/yanzay/tbot/v2"
)

type Invoice struct {
	bot *Bot
	ctx context.Context
}

func NewInvoice(bot *Bot, ctx context.Context) *Invoice {
	return &Invoice{bot: bot, ctx: ctx}
}

func (i *Invoice) SendInvoice(cq *tbot.CallbackQuery) {
	invoiceMiddleware := NewInvoiceMiddleware(i.bot, i.ctx, cq.Message)

	switch callBackMessage := cq.Data; callBackMessage {
	case "getInvoice":
		invoiceMiddleware.SendPayLink(1)
	case "getInvoiceYear":
		invoiceMiddleware.SendPayLink(11)
	}

}
