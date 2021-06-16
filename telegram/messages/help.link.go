package messages

import (
	"fmt"
	"github.com/yanzay/tbot/v2"
)

func LinkButton() *tbot.InlineKeyboardMarkup {
	button := tbot.InlineKeyboardButton{
		Text: "Чат поддержки",
		URL:  "tg://resolve?domain=PredictStock_Help",
	}

	return &tbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]tbot.InlineKeyboardButton{
			{button},
		},
	}
}

func HelpMessage() string {
	return fmt.Sprintf("Если у вас возникли вопросы или нужна помощь, \nперейдите в наш чат поддержки")
}
