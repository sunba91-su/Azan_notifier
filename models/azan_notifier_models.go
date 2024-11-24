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
	Sender   int64    `json:"lineNumber"`
	Message  string   `json:"MessageText"`
	Recivers []string `json:"Mobiles"`
	SendTime *int64   `json:"SendDateTime,omitempty"`
}

type SMSResponse struct {
	Status   int    `json:"status"`
	Response string `json:"message"`
}

type EventUnixTime struct {
	Imsaak     int64
	ImsaakEXP  int64
	Sunrise    int64
	Noon       int64
	NoonEXP    int64
	Sunset     int64
	Maghreb    int64
	MaghrebEXP int64
	Midnight   int64
}

var EventMessages = map[string]string{
	"Imsaak":     "اذان صبح به افق شهر %s",
	"ImsaakEXP":  "یادآوری : \n \t پانزده دقیقه تا قضای نماز صبح به افق شهر %s",
	"Sunrise":    "اعلام طلوع شرعی خورشید به افق شهر %s",
	"Noon":       "اذان ظهر به افق شهر %s",
	"NoonEXP":    "یادآوری : \n \t پانزده دقیقه تا قضای نماز ظهر به افق شهر %s",
	"Sunset":     "اعلام غروب شرعی خورشید به افق شهر %s",
	"Maghreb":    "اذان مغرب به افق شهر %s",
	"MaghrebEXP": "یادآوری : \n \t پانزده دقیقه تا قضای نماز مغرب و عشا به افق شهر %s",
	"Midnight":   "اعلام نیمه شب شرعی به افق شهر %s",
}
