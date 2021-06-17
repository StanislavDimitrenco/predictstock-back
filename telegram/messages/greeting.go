package messages

import (
	"fmt"
	"github.com/yanzay/tbot/v2"
	"time"
)

func Greeting(m *tbot.Message) string {
	return fmt.Sprintf("Hello, %s", m.From.FirstName)
}

func TrialMessage() string {
	return fmt.Sprintf("Your free trial period is activated until %s. Knowledge ➡️ profit!", time.Now().Add(time.Hour*24*7).Format("2006-01-02"))
}
