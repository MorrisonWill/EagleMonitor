package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/MorrisonWill/EagleMonitor/eagleapps"

	t "github.com/MorrisonWill/EagleMonitor/eagleapps/types"
)

func ConstructCourseList() (CourseList []t.CourseStruct) {
	courses := eagleapps.GetCourses("10000", "1")
	fmt.Printf("got %d courses\n", len(courses))
	for i, course := range courses {
		fmt.Printf("finished data on %d\n", i)
		activityOfferings := eagleapps.GetActivityOfferings(course.ReferenceID)
		seatCounts := eagleapps.GetSeatCountsFromOfferings(activityOfferings)

		instructors := []string{}
		for _, offering := range activityOfferings {
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
	eagleapps.Authenticate("morriswk@bc.edu", "***REMOVED***")

	course := ConstructCourseList()

	for i, item := range course {
		fmt.Printf("done %d\n", i)
		description := item.Description.(t.Description)
		descriptionText := description.Description
		removeHtmlPattern := regexp.MustCompile("<[^>]*>")
		descriptionText = removeHtmlPattern.ReplaceAllString(descriptionText, "")

		client := http.Client{}

		data, _ := json.Marshal(map[string]interface{}{
			"id":               item.ID,
			"name":             item.Name,
			"description":      descriptionText,
			"courseLevel":      description.CourseLevel,
			"status":           description.Status,
			"school":           description.School,
			"coreRequirements": description.CoreProgramrequirements,
			"creditCount":      description.CreditCount,
			"coRequisites":     description.CoRequisites,
			"preRequisites":    description.PreRequisites,
			"restrictions":     description.Restrictions,
			"termID":           item.TermID,
			"instructors":      item.Instructors,
			"seatData":         CreateSeatDataStringSlice(item.SeatData),
		})

		request, _ := http.NewRequest("POST", "https://npthxfsjlkgtythiccdg.supabase.co/rest/v1/courses", bytes.NewBuffer(data))
		request.Header.Add("apikey", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYW5vbiIsImlhdCI6MTYzNjU2ODc3NSwiZXhwIjoxOTUyMTQ0Nzc1fQ.CwUnJhI0_fMS5O77DxLKkhHf9ULziq5eCvHwEY9z5RI")
		request.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYW5vbiIsImlhdCI6MTYzNjU2ODc3NSwiZXhwIjoxOTUyMTQ0Nzc1fQ.CwUnJhI0_fMS5O77DxLKkhHf9ULziq5eCvHwEY9z5RI")
		request.Header.Add("Content-Type", "application/json")
		request.Header.Add("Prefer", "return=representation")

		client.Do(request)
	}
}
