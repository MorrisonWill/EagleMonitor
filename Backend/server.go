package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"

	"github.com/MorrisonWill/EagleMonitor/eagleapps"
	database "github.com/MorrisonWill/EagleMonitor/easydb"
)

type Config struct {
	EagleApps struct {
		User string `json:"User"`
		Pass string `json:"Pass"`
	} `json:"EagleApps"`
	Database struct {
		String string `json:"String"`
		Token  string `json:"Token"`
	} `json:"Database"`
	Port      string `json:"Port"`
	JwtSecret string
}

type Courses struct {
	CourseOfferingIds []string `json:"CourseOfferingIds"`
}

type User struct {
	Email   string  `json:"Email"`
	UserId  string  `json:"UserId"`
	Courses Courses `json:"Courses"`
}

var tokenAuth *jwtauth.JWTAuth
var db *database.DB

func router() http.Handler {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
			course := User{
				Email:   "test@gmail.com",
				UserId:  "123",
				Courses: Courses{},
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(course)
		})
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
		r.Post("/watchCourses", func(w http.ResponseWriter, r *http.Request) {
			_, claims, _ := jwtauth.FromContext(r.Context())

			courses := Courses{}
			json.NewDecoder(r.Body).Decode(&courses)

			fmt.Println(courses)

			user := User{
				Email:   fmt.Sprintf("%v", claims["email"]),
				UserId:  fmt.Sprintf("%v", claims["user_id"]),
				Courses: courses,
			}

			db.Put(fmt.Sprintf("%v", claims["user_id"]), user)
			fmt.Println(db.List())
			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(user)
		})
	})

	return r
}

func main() {
	config := Config{}
	configFile, err := os.Open("serverConfig.json")

	defer configFile.Close()

	if err != nil {
		log.Fatalln("Failed to open config file", err)
	}

	json.NewDecoder(configFile).Decode(&config)

	// eagleapps.Authenticate(config.EagleApps.User, config.EagleApps.Pass)
	// db = database.Connect(config.Database.String, config.Database.Token)

	fmt.Printf("Listening on port %s\n", config.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", config.Port), router())
}
