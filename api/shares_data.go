package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type ShareData struct {
	Doing  string  `json:"doing"`
	Rating float32 `json:"rating"`
	Days   int32   `json:"days"`
	Name   string  `json:"name"`
	Date   string  `json:"date"`
	Ticker string  `json:"ticker"`
	Status string  `json:"status"`
}

func (s ShareData) GetDays() int32 {
	return s.Days
}

func (s ShareData) GetStatus() string {
	return s.Status
}

func (s ShareData) GetDate() time.Time {
	layout := "2006-01-02"
	t, _ := time.Parse(layout, s.Date)
	return t
}

func (s ShareData) GetRating() float32 {
	return s.Rating
}

func (s ShareData) GetDoing() string {
	text := strings.ToUpper(s.Doing)
	text = strings.TrimSpace(text)
	if text != "N/A" {
		return s.Doing
	}
	return "Insufficient data to calculate a valid consensus"
}

func (s ShareData) GetTicker() string {
	return s.Ticker
}

//constructor new ShareData
func NewShareData() ShareData {
	var s ShareData
	return s
}

func (s ShareData) GetName() string {
	return s.Name
}

//get response from API request
func GetResponse(ticker string) []byte {

	url := os.Getenv("API_HOST") + "market/ticker/" + ticker

	client := http.Client{
		Timeout: 30 * time.Second,
	}
	response, err := client.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalln(err)
	}

	return body
}

//decoding JSON byte array to ShareData struct
func (s ShareData) Decoder(jsonData []byte) (ShareData, error) {

	err := json.Unmarshal(jsonData, &s)
	if err != nil {
		return s, err
	}

	return s, nil
}
