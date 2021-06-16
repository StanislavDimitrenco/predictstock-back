package messages

import (
	"fmt"
	"time"
)

func ShareNotFound() string {
	return "_No matches found_"
}

func ErrorServer() string {
	return "_Some error_"
}

type ShareData interface {
	GetName() string
	GetTicker() string
	GetDays() int32
	GetStatus() string
	GetDate() time.Time
	GetRating() float32
	GetDoing() string
}

func Share(share ShareData) string {
	var daysText string
	if share.GetDays() == 1 {
		daysText = "day"
	} else {
		daysText = "days"
	}

	return fmt.Sprintf("<b>%s</b>\n"+
		"<b>%s</b>\n"+
		"<b>%s</b>\n"+
		"Consensus: <b>%s</b>\n"+
		"Status: <b>%s %d %s</b> \n",
		share.GetDate().Format("Jan2'06"),
		share.GetTicker(),
		share.GetName(),
		share.GetDoing(),
		share.GetStatus(),
		share.GetDays(),
		daysText,
	)

}
