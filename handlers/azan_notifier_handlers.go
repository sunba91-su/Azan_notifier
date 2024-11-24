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
	parsedTime, err := time.Parse(layout, fullTimeString)
	fmt.Println("the time string: ", fullTimeString)
	fmt.Println("the parset time : ", parsedTime)
	if err != nil {
		return 0
	}
	location, err := time.LoadLocation("Asia/Tehran")
	if err != nil {
		return 0
	}
	timeInLocation, _ := time.ParseInLocation(layout, fullTimeString, location)
	return timeInLocation.Unix()
}

func GenEventsTimes(data models.GetSunsetInfo) models.EventUnixTime {
	var ReligiousUnixTimes models.EventUnixTime
	ReligiousUnixTimes.Imsaak = ParsTime(data.Imsaak)
	ReligiousUnixTimes.Sunrise = ParsTime(data.Sunrise)
	ReligiousUnixTimes.Noon = ParsTime(data.Noon)
	ReligiousUnixTimes.Sunset = ParsTime(data.Sunset)
	ReligiousUnixTimes.Maghreb = ParsTime(data.Maghreb)
	ReligiousUnixTimes.Midnight = ParsTime(data.Midnight)
	return ReligiousUnixTimes
}
