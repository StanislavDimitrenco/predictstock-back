package telegram

import (
	"github.com/Paramosch/predictstock-backend-eng/logger"
	"github.com/yanzay/tbot/v2"
	"strconv"
	"strings"
)

type Message struct {
	text   string
	chatId int
}

func (m Message) Text() string {
	return m.text
}

func (m Message) ChatId() int {
	return m.chatId
}

var Messages = make(chan Message)

func PushMessage(message string, chatId int) {
	go func() { Messages <- Message{message, chatId} }()
}

func Sender(bot *Bot) {
	for {
		select {
		case message := <-Messages:
			text := strings.ReplaceAll(message.Text(), "!", "\\!")
			if _, err := bot.Client.SendMessage(strconv.Itoa(message.ChatId()), text, tbot.OptParseModeMarkdown); err != nil {
				logger.NewLogger("telegram-sender").SetFields(logger.Fields{
					"error":   err.Error(),
					"message": text,
				}).LogError()
			}
		}
	}
}
