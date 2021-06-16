package telegram

import (
	"context"
	"fmt"
	"github.com/Paramosch/predictstock-backend-eng/api"
	"github.com/Paramosch/predictstock-backend-eng/database"
	"github.com/Paramosch/predictstock-backend-eng/logger"
	"github.com/Paramosch/predictstock-backend-eng/telegram/messages"
	"github.com/yanzay/tbot/v2"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func Share(bot *Bot, ctx context.Context) {
	bot.Server.HandleMessage("", func(m *tbot.Message) {
		_ = bot.Client.SendChatAction(m.Chat.ID, tbot.ActionTyping)

		invoiceMiddleware := NewInvoiceMiddleware(bot, ctx, m)
		if !invoiceMiddleware.IsPaid() {
			invoiceMiddleware.SendPayMessage()
			return
		}

		//save user history
		dbUser := database.NewUserRepo(ctx.Value("db").(*gorm.DB))

		intTelegramId, _ := strconv.Atoi(m.Chat.ID)
		telegramID := map[string]interface{}{"telegram_id": intTelegramId}
		user, _ := dbUser.FindBy(telegramID)
		dataHistory := database.NewUserHistoryRepo(ctx.Value("db").(*gorm.DB))

		dataHistory.Create(
			&database.UserHistory{
				UserID:    user.ID,
				Message:   m.Text,
				CreatedAt: time.Now(),
			},
		)

		share, err := api.NewShareData().Decoder(api.GetResponse(m.Text))
		if err != nil {
			logFile := logger.NewLogger("invalid_api_response")
			logFile.LogError(err)
			_, _ = bot.Client.SendMessage(m.Chat.ID, messages.ErrorServer(), tbot.OptParseModeMarkdown)
		} else if blank := api.NewShareData(); share == blank {
			_, _ = bot.Client.SendMessage(m.Chat.ID, messages.ShareNotFound(), tbot.OptParseModeMarkdown)
		} else {
			fmt.Println(messages.Share(share))
			_, err := bot.Client.SendMessage(m.Chat.ID, messages.Share(share), tbot.OptParseModeHTML)
			if err != nil {
				fmt.Println(err)
			}
		}
	})
}
