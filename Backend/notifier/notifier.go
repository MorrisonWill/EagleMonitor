package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/mailgun/mailgun-go"
)

type Profile struct {
	ID      string   `json:"id"`
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

func getProfiles() []Profile {
	request, _ := http.NewRequest("GET", "https://npthxfsjlkgtythiccdg.supabase.co/rest/v1/profiles?select=id,courses,email", nil)

	request.Header.Add("apikey", serviceKey)
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", serviceKey))

	response, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	profiles := &[]Profile{}

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

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}

func dropCourseFromProfile(profile Profile, course string) {
	index := indexOf(course, profile.Courses)
	payload, _ := json.Marshal(append(profile.Courses[:index], profile.Courses[index+1:]...))

	fmt.Printf("removing %s, new %v\n", course, append(profile.Courses[:index], profile.Courses[index+1:]...))

	j := json.RawMessage(payload)

	data, _ := json.Marshal(map[string]interface{}{
		"courses": j,
	})
	request, err := http.NewRequest("PATCH", fmt.Sprintf("https://npthxfsjlkgtythiccdg.supabase.co/rest/v1/profiles?id=eq.%s", profile.ID), bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Add("apikey", serviceKey)
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", serviceKey))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Prefer", "return=representation")

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	time.Sleep(time.Millisecond * 250)
}

func monitorAndSend() {
	profiles := getProfiles()

	for _, profile := range profiles {
		for _, course := range profile.Courses {
			if checkCourseStatus(course) == "open" {
				dropCourseFromProfile(profile, course)
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
