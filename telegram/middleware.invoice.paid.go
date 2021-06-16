package telegram

import (
	"context"
	"fmt"
	"github.com/Paramosch/predictstock-backend-eng/database"
	"github.com/Paramosch/predictstock-backend-eng/logger"
	"github.com/Paramosch/predictstock-backend-eng/payment"
	"github.com/Paramosch/predictstock-backend-eng/telegram/messages"
	"github.com/yanzay/tbot/v2"
	"gorm.io/gorm"
	"os"
	"strconv"
	"time"
)

type InvoiceMiddleware struct {
	bot *Bot
	ctx context.Context
	db  *gorm.DB
	m   *tbot.Message
	qb  string
}

func NewInvoiceMiddleware(bot *Bot, ctx context.Context, m *tbot.Message) *InvoiceMiddleware {
	return &InvoiceMiddleware{bot: bot, ctx: ctx, db: ctx.Value("db").(*gorm.DB), m: m}
}

func (i InvoiceMiddleware) User() *database.User {
	repo := database.NewUserRepo(i.db)
	user, notFound := repo.FindBy(map[string]interface{}{"telegram_id": i.m.Chat.ID})
	if notFound {
		_, _ = i.bot.Client.SendMessage(i.m.Chat.ID, messages.TrialMessage(), tbot.OptParseModeHTML)

		telegramId, _ := strconv.Atoi(i.m.Chat.ID)
		user = repo.Create(&database.User{
			TelegramId: telegramId,
			Username:   i.m.Chat.Username,
			Name:       i.m.Chat.FirstName,
			Lastname:   i.m.Chat.LastName,
			IsPaid:     true,
			IsActive:   true,
			PaidUntil:  time.Now().Add(time.Hour * 24 * 7),
		})
	}
	repo.CheckPaidUntil(user)

	return user
}

func (i InvoiceMiddleware) IsPaid() bool {
	user := i.User()
	if user.GetIsPaid() {
		return true
	} else {
		return false
	}
}

//Edit message. Change button to pay link
func (i InvoiceMiddleware) SendPayLink(mounts int) {
	user := i.User()
	invoiceService := payment.NewInvoice(i.ctx)
	NewInvoiceCleaner(i.ctx, i.bot).Clean(user)

	var message string
	var invoice *database.Invoice
	var url string

	switch mounts {
	case 1:
		price, _ := strconv.ParseInt(os.Getenv("INVOICE_PRICE"), 0, 64)
		fmt.Println(price)
		invoice, url = invoiceService.Create(user.GetId(), mounts, price)
		message = messages.InvoiceMessage(
			os.Getenv("INVOICE_PRICE"),
		)

	case 11:
		price, _ := strconv.ParseInt(os.Getenv("INVOICE_PRICE"), 0, 64)
		price = price * 11
		invoice, url = invoiceService.Create(user.GetId(), mounts+1, price)
		fmt.Println(price)
		message = messages.InvoiceMessage(
			strconv.FormatInt(price, 10),
		)
	}

	sentMessage, err := i.bot.Client.SendMessage(
		i.m.Chat.ID,
		message,
		tbot.OptInlineKeyboardMarkup(messages.BuyPlanLink(url)),
		tbot.OptParseModeMarkdown,
	)
	if err == nil && sentMessage != nil {
		invoiceService.AttachMessageId(invoice, sentMessage.MessageID)
	} else {
		i.logMessageFail(message, err)
	}
}

func (i InvoiceMiddleware) SendPayMessage() {
	message := messages.BuyPlan(
		os.Getenv("INVOICE_PRICE"),
	)
	_, err := i.bot.Client.SendMessage(
		i.m.Chat.ID,
		message,
		tbot.OptParseModeMarkdown,
		tbot.OptInlineKeyboardMarkup(messages.GetInvoiceButton()),
	)
	if err != nil {
		i.logMessageFail(message, err)
	}
}

func (i InvoiceMiddleware) SendPayProlongationLink() {
	user := i.User()
	message := messages.ProlongationPlan(
		user.GetPaidUntil().AddDate(0, 1, 0),
		os.Getenv("INVOICE_PRICE"),
	)

	_, err := i.bot.Client.SendMessage(
		i.m.Chat.ID,
		message,
		tbot.OptParseModeMarkdown,
		tbot.OptInlineKeyboardMarkup(messages.GetInvoiceButton()),
	)

	fmt.Println(err)
	if err != nil {
		i.logMessageFail(message, err)
	}
}

func (i InvoiceMiddleware) logMessageFail(message string, err error) {
	logFile := logger.NewLogger("telegram-sender")
	logFile.SetFields(logger.Fields{"message": message}).LogError(err)
}
