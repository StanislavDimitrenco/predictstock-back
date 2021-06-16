package messages

import (
	"fmt"
	"github.com/Paramosch/predictstock-backend-eng/database"
)

func Profile(user *database.User) string {
	paidStatus := "Состояние тарифа: "
	paidProlongation := ""
	if user.GetIsPaid() {
		paidStatus = paidStatus + "*Оплачен*"
		paidProlongation = fmt.Sprintf(
			"\nДействие вашего тарифа до: *%v*\n",
			user.GetPaidUntil().Format("02-01-2006"),
		)
	} else {
		paidStatus = paidStatus + "*Не оплачен*"
	}

	return fmt.Sprintf("%s\n%s", paidStatus, paidProlongation)
}
