package messages

import (
	"fmt"
	"github.com/yanzay/tbot/v2"
	"strconv"
	"strings"
	"time"
)

func GetInvoiceButton() *tbot.InlineKeyboardMarkup {
	button1 := tbot.InlineKeyboardButton{
		Text:         "Monthly payment",
		CallbackData: "getInvoice",
	}

	button2 := tbot.InlineKeyboardButton{
		Text:         "Annual payment",
		CallbackData: "getInvoiceYear",
	}

	return &tbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]tbot.InlineKeyboardButton{
			{button1},
			{button2},
		},
	}
}

func BuyPlanLink(url string) *tbot.InlineKeyboardMarkup {
	button1 := tbot.InlineKeyboardButton{
		Text: "Pay",
		URL:  url,
	}

	return &tbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]tbot.InlineKeyboardButton{
			{button1},
		},
	}
}

func InvoiceMessage(price string) string {
	return fmt.Sprintf("Invoice for payment of the tariff for the amount\\: *%s USD*", price)
}

func BuyPlan(price string) string {

	pricePerMonth, _ := strconv.Atoi(price)
	pricePerMonth = pricePerMonth * 12
	return fmt.Sprintf(
		"To use the bot, pay for the tariff\\.\n\nThe tariff includes a subscription for 30 days\\.\n\nSubscription cost \\-  %s USD\n\nWhen paying for the year in the amount %d USD, we give you a month for free\n\n",
		price,
		pricePerMonth,
	)
}

func ProlongationPlan(timeProlongation time.Time, price string) string {
	//newTime := strings.ReplaceAll(timeProlongation.Format("2006-01-02"), "-", "\\-")
	newTimeAfterYear := strings.ReplaceAll(timeProlongation.AddDate(0, 11, 0).Format("02-01-2006"), "-", "\\-")

	priceForYear, _ := strconv.Atoi(price)
	priceForYear = priceForYear * 11
	prc := strconv.Itoa(priceForYear)

	return fmt.Sprintf(
		"You can renew your subscription for a month \n"+
			"at just *%s USD*\n\n"+
			"*SPECIAL OFFER\\:*\n"+
			"annual subscription for just *%s USD*\n"+
			"\\- one month OFF\\! \n"+
			"The tariff will last until *%s*\\.",
		price,
		prc,
		newTimeAfterYear,
	)
}
