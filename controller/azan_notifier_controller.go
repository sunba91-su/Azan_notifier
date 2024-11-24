package controller

import (
	"azan_notifier/handlers"
	"azan_notifier/models"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"
)

func StartProgram() {
	CityCode := handlers.GetEnv("CityCode")
	ReligiousTimes, err := GetReligiousTimes(CityCode)
	if err != nil {
		fmt.Println("An error accoure", err)
	}
	DailyReportResponse := handlers.GenDailyReport(ReligiousTimes)
	ReligiousUnixTimes := handlers.GenEventsTimes(ReligiousTimes)
	ScheduleEventNotif(ReligiousUnixTimes, ReligiousTimes.City)
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
	DailyReportBody := handlers.ReqBodyGenerator(handlers.GetResivers(), Message, nil)
	//TODO Remove Debug Log
	fmt.Println("****************", DailyReportBody, "**************************")
	// handlers.SendSMS(DailyReportBody)
}
func ScheduleEventNotif(ReligiousTime models.EventUnixTime, City string) {
	EventList := reflect.ValueOf(ReligiousTime)
	for Event := 0; Event < EventList.NumField(); Event++ {
		field := EventList.Type().Field(Event)
		EventTime := EventList.Field(Event).Int()
		EventSMSBody := GenerateScheduleEventNotifBody(field.Name, EventTime, City)
		if EventTime < time.Now().Unix() {
			fmt.Println(EventSMSBody)
			// handlers.SendSMS(EventSMSBody)
		}
	}
}

func GenerateScheduleEventNotifBody(Event string, EventUnixTime int64, City string) models.SendSMS {
	var EventSMSModel models.SendSMS
	EventMessage, _ := handlers.GetEventMessage(Event)
	EventMessages := fmt.Sprintf(EventMessage, City)
	EventSMSModel = handlers.ReqBodyGenerator(handlers.GetResivers(), EventMessages, &EventUnixTime)
	return EventSMSModel
}
