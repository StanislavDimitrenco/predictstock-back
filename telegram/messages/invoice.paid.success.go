package messages

import (
	"fmt"
)

func InvoicePaidSuccess(invoice int64, expiredAt string) string {
	return fmt.Sprintf("Ваш счёт *№%d* был успешно оплачен\\.\nТариф продлён до *%s*\\.\n\n", invoice, expiredAt)
}
