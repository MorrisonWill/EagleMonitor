package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/mailgun/mailgun-go"
)

type Profiles []struct {
	Email   string   `json:"email"`
	Courses []string `json:"courses"`
}

type Status []struct {
	Status string `json:"status"`
}

var domain = "mail.eaglemonitor.app"
var privateKey string // will be set later from env

var serviceKey string // to be set later from env

var client = http.Client{}

func sendNotifyEmail(email, course string) (string, error) {
	mg := mailgun.NewMailgun(domain, privateKey)
	sender := "notification@eaglemonitor.app"
	subject := fmt.Sprintf("%s IS OPEN!", course)
	body := fmt.Sprintf("One of the courses you added to your monitoring list on Eagle Monitor, %s, has an opening! Grab it while you can!", course)

	message := mg.NewMessage(sender, subject, body, email)

	resp, _, err := mg.Send(message)

	if err != nil {
		log.Fatalln(err)
	}

	return resp, err
}

func getProfiles() Profiles {
	request, _ := http.NewRequest("GET", "https://npthxfsjlkgtythiccdg.supabase.co/rest/v1/profiles?select=courses,email", nil)

	request.Header.Add("apikey", serviceKey)
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", serviceKey))

	response, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	profiles := &Profiles{}

	json.NewDecoder(response.Body).Decode(profiles)

	return *profiles
}

func checkCourseStatus(course string) string {
	request, _ := http.NewRequest("GET", fmt.Sprintf("https://npthxfsjlkgtythiccdg.supabase.co/rest/v1/courses?id=eq.%s&select=status", course), nil)

	request.Header.Add("apikey", serviceKey)
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", serviceKey))

	response, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	status := &Status{}

	json.NewDecoder(response.Body).Decode(status)

	return (*status)[0].Status
}

func monitorAndSend() {
	profiles := getProfiles()

	for _, profile := range profiles {
		for _, course := range profile.Courses {
			if checkCourseStatus(course) == "open" {
				sendNotifyEmail(profile.Email, course)
			}
		}
	}
}

func main() {
	godotenv.Load(".env")

	serviceKey = os.Getenv("SUPABASE_SERVICE_KEY")
	privateKey = os.Getenv("MAILGUN_PRIVATE_KEY")

	monitorAndSend()
}
