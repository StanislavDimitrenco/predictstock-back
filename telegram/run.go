package telegram

import (
	"context"
	"github.com/yanzay/tbot/v2"
	"log"
	"os"
)

type Bot struct {
	Client tbot.Client
	Server tbot.Server
}

func NewBot() *Bot {
	server := tbot.New(os.Getenv("TELEGRAM_TOKEN"))
	return &Bot{Client: *server.Client(), Server: *server}
}

func Run(bot *Bot, ctx context.Context) *Bot {

	//Run messages pusher
	go Sender(bot)

	//Define commands
	Define(bot, ctx)

	log.Fatal(bot.Server.Start())

	return bot
}
