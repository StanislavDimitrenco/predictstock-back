package telegram

import (
	"context"
	"fmt"
	"github.com/Paramosch/predictstock-backend-eng/telegram/messages"
	"github.com/yanzay/tbot/v2"
	"strings"
)

var info = tbot.KeyboardButton{
	Text:            "⚙️ How to use",
	RequestContact:  false,
	RequestLocation: false,
	RequestPoll:     nil,
}

var profile = tbot.KeyboardButton{
	Text:            "💳 Profile",
	RequestContact:  false,
	RequestLocation: false,
	RequestPoll:     nil,
}

var help = tbot.KeyboardButton{
	Text:            "❓ Support",
	RequestContact:  false,
	RequestLocation: false,
	RequestPoll:     nil,
}

var about = tbot.KeyboardButton{
	Text:            "ℹ️ Info",
	RequestContact:  false,
	RequestLocation: false,
	RequestPoll:     nil,
}

func MenuButtons() *tbot.ReplyKeyboardMarkup {
	return &tbot.ReplyKeyboardMarkup{
		Keyboard: [][]tbot.KeyboardButton{
			{info, profile},
			{help, about},
		},
		ResizeKeyboard:  true,
		OneTimeKeyboard: false,
	}
}

func SetupMenu(bot *Bot, ctx context.Context) {
	bot.Server.HandleMessage(info.Text, func(m *tbot.Message) {
		_, err := bot.Client.SendMessage(m.Chat.ID, messages.Tutorial(), tbot.OptParseModeMarkdown)
		fmt.Println(err)
	})

	bot.Server.HandleMessage(help.Text, func(m *tbot.Message) {
		_, err := bot.Client.SendMessage(m.Chat.ID, messages.HelpMessage(), tbot.OptInlineKeyboardMarkup(messages.LinkButton()))

		if err != nil {
			fmt.Println(err)
		}
	})

	bot.Server.HandleMessage(about.Text, func(m *tbot.Message) {
		_, err := bot.Client.SendMessage(m.Chat.ID, messages.AboutMessage(), tbot.OptParseModeMarkdown)
		fmt.Println(err)
	})

	bot.Server.HandleMessage(profile.Text, func(m *tbot.Message) {
		invoiceMiddleware := NewInvoiceMiddleware(bot, ctx, m)
		mess := strings.ReplaceAll(messages.Profile(invoiceMiddleware.User()), "-", "\\-")
		_, err := bot.Client.SendMessage(m.Chat.ID, mess, tbot.OptParseModeMarkdown)
		if err != nil {
			fmt.Println(err)
			fmt.Println(mess)
		}
		if !invoiceMiddleware.IsPaid() {
			invoiceMiddleware.SendPayMessage()
		} else {
			invoiceMiddleware.SendPayProlongationLink()
		}
	})
}
