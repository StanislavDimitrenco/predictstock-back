package telegram

import (
	"context"
	"github.com/Paramosch/predictstock-backend-eng/database"
	"github.com/yanzay/tbot/v2"
	"gorm.io/gorm"
	"os"
	"strconv"
)

//UpdateMenu update all buttons
func UpdateMenu(ctx context.Context, bot *Bot) {
	if os.Getenv("UPDATE_MENU") == "yes" {
		db := ctx.Value("db").(*gorm.DB)
		users := database.NewUserRepo(db).GetAll()

		for _, user := range users {
			_, _ = bot.Client.SendMessage(strconv.Itoa(user.TelegramId), "Бот обновился. Обратите внимание на новый функционал.", tbot.OptReplyKeyboardMarkup(MenuButtons()))
		}
	}
}
