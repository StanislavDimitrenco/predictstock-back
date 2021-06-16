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
		Text:         "Запросить счёт на месяц",
		CallbackData: "getInvoice",
	}

	button2 := tbot.InlineKeyboardButton{
		Text:         "Запросить счёт на год",
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
		Text: "Оплатить",
		URL:  url,
	}

	return &tbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]tbot.InlineKeyboardButton{
			{button1},
		},
	}
}

func InvoiceMessage(price string) string {
	return fmt.Sprintf("Счёт на оплату тарифа на сумму\\: *%sр*", price)
}

func BuyPlan(price string) string {

	pricePerMonth, _ := strconv.Atoi(price)
	pricePerMonth = pricePerMonth * 12
	return fmt.Sprintf(
		"Для использования бота оплатите тариф\\.\n\nВ тариф входит подписка на 30 дней\\.\n\nСтоимость подписки \\-  %s₽\n\nПри оплатае за год суммой %d₽, мы дарим вам месяц бесплатно\n\n",
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
		"Вы можете продлить тариф на месяц\\. \n"+
			"Стоимость продления \\- *%s₽*\n\n"+
			"*БОНУС\\!*\n"+
			"При оплате тарифа за год \\- *один месяц бесплатно\\!*\n"+
			"Стоимость продления \\- *%s₽*\n"+
			"Тариф продлится до *%s*\\.",
		price,
		prc,
		newTimeAfterYear,
	)
}
