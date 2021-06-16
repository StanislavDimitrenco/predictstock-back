package telegram

import (
	"context"
)

func Define(bot *Bot, ctx context.Context) {
	Start(bot, ctx)
	SetupMenu(bot, ctx)
	Share(bot, ctx)

	bot.Server.HandleCallback(NewInvoice(bot, ctx).SendInvoice)
}
