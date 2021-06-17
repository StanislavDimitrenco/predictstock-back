package messages

import (
	"fmt"
	"github.com/Paramosch/predictstock-backend-eng/database"
)

func Profile(user *database.User) string {
	paidStatus := "Tariff status: "
	paidProlongation := ""
	if user.GetIsPaid() {
		paidStatus = paidStatus + "*Paid Up*"
		paidProlongation = fmt.Sprintf(
			"\nYour subscription is paid till: *%v*\n",
			user.GetPaidUntil().Format("Jan 2, 2006"),
		)
	} else {
		paidStatus = paidStatus + "*Not Paid*"
	}

	return fmt.Sprintf("%s\n%s", paidStatus, paidProlongation)
}
