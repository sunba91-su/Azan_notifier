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
		"سلام عزیزم \n روز به خیر \n اطلاعات روز : \n تاریخ : %s \n اذان صبح : %s \n طلوع خورشید : %s \n اذان ظهر : %s \n غروب خورشید : %s \n اذان مغرب و عشا : %s \n نیمه شب شرعی : %s",
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
	fmt.Println(MessageBody)
	MessageByte, _ := json.Marshal(MessageBody)
	fmt.Println(MessageByte)
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
