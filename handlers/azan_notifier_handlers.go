package handlers

import (
	"azan_notifier/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

func LoadEnvs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func GetEnv(env string) string {
	return os.Getenv(env)
}
func GetIntEnv(env string) int64 {
	myIntegerStr := GetEnv(env)
	myInteger, err := strconv.ParseInt(myIntegerStr, 10, 64)
	if err != nil {
		log.Fatalf("Error converting MY_INTEGER to int: %v", err)
	}
	return myInteger
}
func GenDailyReport(data models.GetSunsetInfo) string {
	DailyReportMessage := fmt.Sprintf(
		"سلام عزیزم \n روز به خیر \n اطلاعات روز شهر %s : \n تاریخ : %s \n اذان صبح : %s \n طلوع خورشید : %s \n اذان ظهر : %s \n غروب خورشید : %s \n اذان مغرب و عشا : %s \n نیمه شب شرعی : %s",
		data.City,
		data.Date,
		data.Imsaak,
		data.Sunrise,
		data.Noon,
		data.Sunset,
		data.Maghreb,
		data.Midnight,
	)
	return DailyReportMessage
}
func SendSMS(MessageBody models.SendSMS) {
	SMSAPI := GetEnv("SMSProviderAPI")
	APIKeyHeader := GetEnv("APIKeyHeader")
	APIKey := GetEnv("APIKey")
	SMSProviderAPIMethod := GetEnv("SMSProviderAPIMethod")
	MessageByte, _ := json.Marshal(MessageBody)
	client := &http.Client{}
	req, err := http.NewRequest(SMSProviderAPIMethod, SMSAPI, bytes.NewBuffer(MessageByte))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add(APIKeyHeader, APIKey)
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
func ParsTime(timeString string) int64 {
	now := time.Now()
	currentDate := now.Format("2006-01-02")
	layout := "2006-01-02 15:04:05"
	fullTimeString := currentDate + " " + timeString
	location, err := time.LoadLocation("Asia/Tehran")
	if err != nil {
		return 0
	}
	timeInLocation, _ := time.ParseInLocation(layout, fullTimeString, location)
	return timeInLocation.Unix()
}
func ReqBodyGenerator(Resivers []string, Message string, SendDate *int64) models.SendSMS {
	SenderNumber := GetIntEnv("Sender")
	Body := models.SendSMS{
		Sender:   SenderNumber,
		Recivers: Resivers,
		Message:  Message,
		SendTime: SendDate,
	}
	return Body
}
func GetResivers() []string {
	Resivers := GetEnv("Resivers")
	Resiver := strings.Split(Resivers, ",")
	return Resiver
}
func GenEventsTimes(data models.GetSunsetInfo) models.EventUnixTime {
	offset := GetIntEnv("EXPTimeOffset")
	var ReligiousUnixTimes models.EventUnixTime
	ReligiousUnixTimes.Imsaak = ParsTime(data.Imsaak)
	ReligiousUnixTimes.Sunrise = ParsTime(data.Sunrise)
	ReligiousUnixTimes.ImsaakEXP = ReligiousUnixTimes.Sunrise - offset
	ReligiousUnixTimes.Noon = ParsTime(data.Noon)
	ReligiousUnixTimes.Sunset = ParsTime(data.Sunset)
	ReligiousUnixTimes.NoonEXP = ReligiousUnixTimes.Sunset - offset
	ReligiousUnixTimes.Maghreb = ParsTime(data.Maghreb)
	ReligiousUnixTimes.Midnight = ParsTime(data.Midnight)
	ReligiousUnixTimes.MaghrebEXP = ReligiousUnixTimes.Midnight - offset
	return ReligiousUnixTimes
}
func DebugSMSBudy(SMSBody models.SendSMS) {
	val := reflect.ValueOf(SMSBody)
	typ := reflect.TypeOf(SMSBody)

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)

		fmt.Printf("************* Debug SMS Body ***************** %s: %v\n", field.Name, value.Interface())
	}
}
func GetEventMessage(Event string) (string, bool) {
	message, exists := models.EventMessages[Event]
	return message, exists
}
