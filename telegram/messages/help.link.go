package messages

import (
	"fmt"
	"github.com/yanzay/tbot/v2"
)

func LinkButton() *tbot.InlineKeyboardMarkup {
	button := tbot.InlineKeyboardButton{
		Text: "Support Chat",
		URL:  "tg://resolve?domain=PredictStock_Help",
	}

	return &tbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]tbot.InlineKeyboardButton{
			{button},
		},
	}
}

func HelpMessage() string {
	return fmt.Sprintf("If you have questions or need help, \nvisit our support chat\\.")
}
