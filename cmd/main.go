package main

import (
	"context"
	"fmt"
	"github.com/Paramosch/predictstock-backend-eng/logger"
	"github.com/Paramosch/predictstock-backend-eng/providers"
	"github.com/Paramosch/predictstock-backend-eng/server"
	"github.com/Paramosch/predictstock-backend-eng/telegram"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"runtime"
)

var ctx context.Context

func main() {

	//listen terminal signal
	terminalOsC := make(chan os.Signal)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Boot all dependencies
	ctx = providers.Boot(context.Background())

	//Handle panic before exit
	//todo разобраться в почему из-за этого бот при остановки заново стартует
	//exitStatus, err := panicwrap.BasicWrap(panicHandler)
	//if err != nil {
	//	panic(err)
	//}
	//if exitStatus >= 0 {
	//	os.Exit(exitStatus)
	//}

	go telegram.Run(ctx.Value("telegramBot").(*telegram.Bot), ctx)

	go telegram.UpdateMenu(ctx, ctx.Value("telegramBot").(*telegram.Bot))

	//turn off the parser
	//go file_parser.Parser(database.NewSharesRepo(ctx.Value("db").(*gorm.DB)))

	_ = server.Run(ctx)
	terminationListening(ctx, terminalOsC)
	fmt.Printf("gorutine - %d", runtime.NumGoroutine())
}

func panicHandler(output string) {
	adminId := os.Getenv("ADMIN_ID")
	bot := ctx.Value("telegramBot").(*telegram.Bot)
	logFile := logger.NewLogger("fatal")
	logFile.LogError(output)
	_, err := bot.Client.SendMessage(adminId, "Bot crash.")
	if err != nil {
		logFile.LogError(err)
	}
	os.Exit(1)
}

func terminationListening(ctx context.Context, terminalOsC chan os.Signal) {
	<-terminalOsC
	// stop web-server
	if webserver, ok := ctx.Value("webserver").(*fiber.App); ok {
		if err := webserver.Shutdown(); err != nil {
			fmt.Println("Can't shutdown server", err)
		} else {
			fmt.Println("Web-Server was successfully stopped")
		}
	}

	os.Exit(1)

}
