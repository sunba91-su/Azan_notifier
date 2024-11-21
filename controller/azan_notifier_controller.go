package controller

import (
	"azan_notifier/handlers"
	"azan_notifier/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func StartProgram() {
	CityCode := handlers.GetEnv("CityCode")
	ReligiousTimes, err := GetReligiousTimes(CityCode)
	if err != nil {
		fmt.Println("An error accoure", err)
	}
	DailyReportResponse := handlers.GenDailyReport(ReligiousTimes)
	DailyReport(DailyReportResponse)

}

func GetReligiousTimes(CityCode string) (models.GetSunsetInfo, error) {
	var response models.GetSunsetInfo
	PrayerAPI := handlers.GetEnv("PrayerAPI")
	CallUrl := fmt.Sprintf("%s/%s", PrayerAPI, CityCode)
	RTimes, err := http.Get(CallUrl)
	if err != nil {
		return response, err
	}
	defer RTimes.Body.Close()

	if RTimes.StatusCode != http.StatusOK {
		return response, fmt.Errorf("failed to get data: %s", RTimes.Status)
	}
	err = json.NewDecoder(RTimes.Body).Decode(&response)
	if err != nil {
		return response, err
	}
	return response, err
}

func DailyReport(Message string) {
	SenderNumber, _ := strconv.Atoi(handlers.GetEnv("Sender"))
	Resivers := handlers.GetEnv("Resivers")
	Resiver := strings.Split(Resivers, ",")
	DailyReportBody := models.SendSMS{
		Sender:   SenderNumber,
		Recivers: Resiver,
		Message:  Message,
	}
	handlers.SendSMS(DailyReportBody)
}
