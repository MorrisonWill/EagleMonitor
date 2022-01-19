package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/MorrisonWill/EagleMonitor/eagleapps"
	"github.com/joho/godotenv"

	t "github.com/MorrisonWill/EagleMonitor/eagleapps/types"
)

var serviceKey string

func ConstructCourseList() (CourseList []t.CourseStruct) {
	courses := eagleapps.GetCourses("10000", "1")
	fmt.Printf("got %d courses\n", len(courses))
	for i, course := range courses {
		fmt.Printf("finished data on %d\n", i)
		activityOfferings := eagleapps.GetActivityOfferings(course.ReferenceID)
		seatCounts := eagleapps.GetSeatCountsFromOfferings(activityOfferings)

		instructors := []string{}
		rooms := []string{}
		times := []string{}

		for _, offering := range activityOfferings {
			if len(offering.ScheduleIds) == 0 {
				continue
			}

			eagleResponse := eagleapps.GetSchedule(offering.ScheduleIds[0])

			if len(eagleResponse) == 0 {
				continue
			}

			scheduleData := eagleResponse[0].ScheduleComponentDisplays
			if len(scheduleData) == 0 {
				continue
			}
			rooms = append(rooms, scheduleData[0].Room.Name)
			if len(scheduleData[0].TimeSlots) == 0 {
				continue
			}
			times = append(times, scheduleData[0].TimeSlots[0].Descr.Plain)

			for _, instructor := range offering.Instructors {
				instructors = append(instructors, instructor.PersonName)
			}
		}

		courseStruct := t.CourseStruct{
			ID:          course.Attributes[0].Value,
			Name:        course.Name,
			Description: course.Descr.Plain,
			TermID:      activityOfferings[0].TermCode,
			Instructors: instructors,
			SeatData:    seatCounts,
			Rooms:       rooms,
			Times:       times,
		}

		CourseList = append(CourseList, courseStruct)
	}
	return
}

func CreateSeatDataStringSlice(s []t.SeatData) (returnSlice []string) {
	for _, data := range s {
		returnSlice = append(returnSlice, fmt.Sprintf("%s:%d:%d:%d", data.Name, data.Total, data.Used, data.Available))
	}
	return
}

func main() {
	godotenv.Load("/home/user/dev/projects/EagleMonitor/Backend/.env")

	serviceKey = os.Getenv("SUPABASE_SERVICE_KEY")

	eagleapps.Authenticate("morriswk@bc.edu", "***REMOVED***")

	course := ConstructCourseList()

	for i, item := range course {
		fmt.Printf("done %d\n", i)
		description := item.Description.(t.Description)
		descriptionText := description.Description
		removeHtmlPattern := regexp.MustCompile("<[^>]*>")
		descriptionText = removeHtmlPattern.ReplaceAllString(descriptionText, "")

		client := http.Client{}

		t := time.Now().Format("01/02/2006 3:4:5 pm")

		fmt.Println(t)

		status := "closed"

		for _, seatStruct := range item.SeatData {
			if seatStruct.Used < seatStruct.Total {
				fmt.Printf("used: %d, total: %d, open\n", seatStruct.Used, seatStruct.Total)
				status = "open"
			}
		}

		data, _ := json.Marshal(map[string]interface{}{
			"id":               item.ID,
			"name":             item.Name,
			"description":      descriptionText,
			"courseLevel":      description.CourseLevel,
			"status":           status,
			"school":           description.School,
			"coreRequirements": description.CoreProgramrequirements,
			"creditCount":      description.CreditCount,
			"coRequisites":     description.CoRequisites,
			"preRequisites":    description.PreRequisites,
			"restrictions":     description.Restrictions,
			"termID":           item.TermID,
			"instructors":      item.Instructors,
			"seatData":         CreateSeatDataStringSlice(item.SeatData),
			"rooms":            item.Rooms,
			"times":            item.Times,
			"lastUpdated":      t,
		})

		// Initial populate request
		// request, _ := http.NewRequest("POST", "https://npthxfsjlkgtythiccdg.supabase.co/rest/v1/courses", bytes.NewBuffer(data))

		// Update request
		request, _ := http.NewRequest("PATCH", fmt.Sprintf("https://npthxfsjlkgtythiccdg.supabase.co/rest/v1/courses?id=eq.%s", item.ID), bytes.NewBuffer(data))

		request.Header.Add("apikey", serviceKey)
		request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", serviceKey))
		request.Header.Add("Content-Type", "application/json")
		request.Header.Add("Prefer", "return=representation")

		response, _ := client.Do(request)

		full, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(full))

		if response.StatusCode != 200 {
			request, _ := http.NewRequest("POST", "https://npthxfsjlkgtythiccdg.supabase.co/rest/v1/courses", bytes.NewBuffer(data))

			request.Header.Add("apikey", serviceKey)
			request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", serviceKey))
			request.Header.Add("Content-Type", "application/json")
			request.Header.Add("Prefer", "return=representation")

			client.Do(request)
		}
	}
}
