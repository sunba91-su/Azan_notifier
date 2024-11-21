package main

import (
	"azan_notifier/controller"
	"azan_notifier/handlers"
)

// func getSunriseSunset(CityCode string) (string, string, error) {
// 	resp, err := http.Get(fmt.Sprintf("%s?lat=%s&lng=%s&formatted=0", PrayerAPI, cityLat(city), cityLng(city)))
// 	if err != nil {
// 		return "", "", err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return "", "", fmt.Errorf("failed to get data: %s", resp.Status)
// 	}

// 	var response SunriseSunsetResponse
// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 		return "", "", err
// 	}

// 	return response.Results.Sunrise, response.Results.Sunset, nil
// }

// func cityLat(city string) string {
// 	// Replace with actual logic to get latitude based on city
// 	return "34.0522" // Example: Los Angeles latitude
// }

// func cityLng(city string) string {
// 	// Replace with actual logic to get longitude based on city
// 	return "-118.2437" // Example: Los Angeles longitude
// }

// func sendSMS(to string, message string) error {
// 	smsPayload := map[string]string{
// 		"to":      to,
// 		"message": message,
// 	}
// 	payloadBytes, _ := json.Marshal(smsPayload)

// 	resp, err := http.Post(smsAPI, "application/json", bytes.NewBuffer(payloadBytes))
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return fmt.Errorf("failed to send SMS: %s", resp.Status)
// 	}

// 	return nil
// }

// func scheduleSMS(sunset string, sunrise string, phoneNumber string) {
// 	sunsetTime, _ := time.Parse(time.RFC3339, sunset)
// 	sunriseTime, _ := time.Parse(time.RFC3339, sunrise)

// 	// Schedule SMS for sunset
// 	time.AfterFunc(sunsetTime.Sub(time.Now()), func() {
// 		sendSMS(phoneNumber, "Sunset is at: "+sunset)
// 	})

// 	// Schedule SMS for sunrise
// 	time.AfterFunc(sunriseTime.Sub(time.Now()), func() {
// 		sendSMS(phoneNumber, "Sunrise is at: "+sunrise)
// 	})
// }f

func main() {
	handlers.LoadEnvs()
	controller.StartProgram()
	// CityCode := handlers.GetEnv("CityCode")
	// PrayerAPI := handlers.GetEnv("PrayerAPI")
	// sunset, sunrise, err := getSunriseSunset(CityCode)
	// if err != nil {
	// 	fmt.Println("Error getting sunrise/sunset:", err)
	// 	return
	// }

	// fmt.Println("Sunrise:", sunrise)
	// fmt.Println("Sunset:", sunset)

	// scheduleSMS(sunset, sunrise, phoneNumber)

	// Keep the program running
	select {}
}
