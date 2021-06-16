package telegram

import (
	"context"
	"github.com/Paramosch/predictstock-backend-eng/telegram/messages"
	"github.com/yanzay/tbot/v2"
)

func Start(bot *Bot, ctx context.Context) {
	bot.Server.HandleMessage("/start", func(m *tbot.Message) {

		_ = bot.Client.SendChatAction(m.Chat.ID, tbot.ActionTyping)

		_, _ = bot.Client.SendMessage(m.Chat.ID, messages.Greeting(m), tbot.OptReplyKeyboardMarkup(MenuButtons()))

		invoiceMiddleware := NewInvoiceMiddleware(bot, ctx, m)
		if invoiceMiddleware.IsPaid() {
			_, _ = bot.Client.SendMessage(m.Chat.ID, messages.Tutorial(), tbot.OptParseModeMarkdown)
		} else {
			invoiceMiddleware.SendPayMessage()
		}
	})
}
