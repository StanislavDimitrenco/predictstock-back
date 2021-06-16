package messages

import (
	"fmt"
	"github.com/yanzay/tbot/v2"
	"time"
)

func Greeting(m *tbot.Message) string {
	return fmt.Sprintf("Приветствуем вас, %s", m.From.FirstName)
}

func TrialMessage() string {
	return fmt.Sprintf("Ваш бесплатный пробный период активирован до %s. Знание ➡️ профит!", time.Now().Add(time.Hour*24*7).Format("2006-01-02"))
}
