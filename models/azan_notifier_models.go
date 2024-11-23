package models

type GetSunsetInfo struct {
	City     string `json:"CityName"`
	TimeZone string `json:"TimeZone"`
	Noon     string `json:"Noon"`
	Sunrise  string `json:"Sunrise"`
	Sunset   string `json:"Sunset"`
	Maghreb  string `json:"Maghreb"`
	Imsaak   string `json:"Imsaak"`
	Date     string `json:"Today"`
	Midnight string `json:"Midnight"`
}

type SendSMS struct {
	Sender   int      `json:"lineNumber"`
	Message  string   `json:"MessageText"`
	Recivers []string `json:"Mobiles"`
	SendTime *int64   `json:"SendDateTime,omitempty"`
}

type SMSResponse struct {
	Status   int    `json:"status"`
	Response string `json:"message"`
}

type EventUnixTime struct {
	Imsaak   int64
	Sunrise  int64
	Noon     int64
	Sunset   int64
	Maghreb  int64
	Midnight int64
}
