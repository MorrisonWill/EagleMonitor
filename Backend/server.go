package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"

	"github.com/MorrisonWill/EagleMonitor/eagleapps"
	database "github.com/MorrisonWill/EagleMonitor/easydb"
)

var tokenAuth *jwtauth.JWTAuth
var db *database.DB

func router() http.Handler {
	fmt.Println("listening")
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		// r.Use(jwtauth.Verifier(tokenAuth))
		// r.Use(jwtauth.Authenticator)
		r.Get("/user/info", func(w http.ResponseWriter, r *http.Request) {
			_, claims, _ := jwtauth.FromContext(r.Context())
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(claims)
		})
		r.Get("/courses/{maxValues}-{startAt}", func(w http.ResponseWriter, r *http.Request) {
			maxValues := chi.URLParam(r, "maxValues")
			startAt := chi.URLParam(r, "startAt")

			courseList := eagleapps.GetCourses(maxValues, startAt)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(courseList)
		})
		r.Get("/activityOfferings/{courseOfferingId}", func(w http.ResponseWriter, r *http.Request) {
			activityOfferings := eagleapps.GetActivityOfferings(chi.URLParam(r, "courseOfferingId"))

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(activityOfferings)
		})
		r.Get("/seatCount/{activityOfferingId}", func(w http.ResponseWriter, r *http.Request) {
			seatCount := eagleapps.GetSeatCount(chi.URLParam(r, "activityOfferingId"))

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(seatCount)
		})
		r.Get("/db", func(w http.ResponseWriter, r *http.Request) {

		})
	})

	return r
}

func main() {
	eaUser := os.Getenv("EA_USER")
	eaPass := os.Getenv("EA_PASS")

	dbStr := os.Getenv("DB_STR")
	dbToken := os.Getenv("DB_TOKEN")

	eagleapps.Authenticate(eaUser, eaPass)
	db = database.Connect(dbStr, dbToken)
	http.ListenAndServe(":8080", router())
}
