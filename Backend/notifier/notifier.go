package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

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

func sendNotifyEmail(email string, courses []string) (string, error) {
	fmt.Printf("notifying %s about %v\n", email, courses)
	mg := mailgun.NewMailgun(domain, privateKey)
	sender := "notification@eaglemonitor.app"
	subject := fmt.Sprintf("%d BC COURSES OPENED!", len(courses))
	list := "'" + strings.Join(courses, `','`) + `'`
	body := fmt.Sprintf("%d of the courses you added to your monitoring list on Eagle Monitor have opened! Here's the list:\n\n%v", len(courses), list)

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

	profiles := []Profile{}

	json.NewDecoder(response.Body).Decode(&profiles)

	return profiles
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

	status := Status{}

	json.NewDecoder(response.Body).Decode(&status)

	return (status)[0].Status
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func replaceCourses(profile Profile, courses []string) {
	payload, _ := json.Marshal(courses)

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
}

func removeAllOpen(courses []string) (open []string, closed []string) {
	for _, course := range courses {
		if checkCourseStatus(course) == "open" {
			open = append(open, course)
		} else {
			closed = append(closed, course)
		}
	}
	return
}

func MonitorAndSend() {
	profiles := getProfiles()

	for _, profile := range profiles {
		open, closed := removeAllOpen(profile.Courses)
		if len(open) > 0 {
			replaceCourses(profile, closed)
			sendNotifyEmail(profile.Email, open)
		}
	}
}

func main() {
	godotenv.Load(".env")

	serviceKey = os.Getenv("SUPABASE_SERVICE_KEY")
	privateKey = os.Getenv("MAILGUN_PRIVATE_KEY")

	test, err := sendNotifyEmail("morriswk@bc.edu", []string{"course goes here"})

	fmt.Println(test)
	fmt.Println(err)

	//	MonitorAndSend()
}
