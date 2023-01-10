package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"github.com/gorilla/mux"
)

// Model for course/author - separate file
type Course struct {
	CourseID    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// Fake DB
var courses []Course

// Middlewares , Helpers - separate file
func (c *Course) IsEmtpty() bool {
	// return c.CourseID == "" && c.CourseName == ""
	return c.CourseName == ""
}

// Controllers = separate file
// Controllers are collection of methods to handle requested routes for particular data resource.
// Whenever a client requests the route, the action performs the business logic code and sends back the response.
// Always pass on a Response and Request object

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	// Write writes the data to the connection as part of an HTTP reply.
	w.Write([]byte("<h1>Welocme to building API</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all the courses")
	// Header returns the header map that will be sent by WriteHeader
	w.Header().Set("Content-Type", "application/json")
	// NewEncoder returns a new encoder that writes to w
	// Encode writes the JSON encoding of v to the stream
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course based on ID")
	w.Header().Set("Content-Type", "application/json")
	// Vars returns the route variables for the current request, if any
	params := mux.Vars(r)
	// Loop through courses, find matching id and return response
	for _, course := range courses {
		if course.CourseID == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// encoder after sending the response automatically exits the function
	json.NewEncoder(w).Encode("No course found with given ID")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add one course into DB")
	w.Header().Set("Content-Type", "application/json")
	// case for empty body
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}
	// what about - {}
	var course Course
	// NewDecoder returns a new decoder that reads from r
	// JSON-encoded is stored in the value pointed to by v
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmtpty() {
		json.NewEncoder(w).Encode("No data inside the JSON")
		return
	}

	// check if the course is duplicate 

	// generate unique string id and append the new course
	rand.Seed(time.Now().UnixNano())
	course.CourseID = strconv.Itoa(rand.Intn(10000))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update course")
	w.Header().Set("Content-Type", "application/json")
	// grab id from the request
	params := mux.Vars(r)
	// loop, get id, remove, add with same id
	for idx, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:idx], courses[idx+1:]...)
			var newCourse Course
			_ = json.NewDecoder(r.Body).Decode(&newCourse)
			newCourse.CourseID = params["id"]
			courses = append(courses, newCourse)
			json.NewEncoder(w).Encode("Updation Successfully")
		}
	}

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course from DB")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for idx, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:idx], courses[idx+1:]...)
			json.NewEncoder(w).Encode("Successfully Deleted the Value")
			break
		}
	}
}

func main() {
	fmt.Println("Building API")
	r := mux.NewRouter()

	// seeding database
	courses = append(courses, Course{CourseID: "1", CourseName: "ReactJS", CoursePrice: 499, Author: &Author{Fullname: "Rajesh", Website: "lco.dev"}})
	courses = append(courses, Course{CourseID: "2", CourseName: "NodeJS", CoursePrice: 299, Author: &Author{Fullname: "Jignesh", Website: "lco.dev"}})

	// routes mapping URL paths to handlers
	// If an incoming request URL matches one of the paths, the corresponding handler is called passing (http.ResponseWriter, *http.Request) as parameters
	r.HandleFunc("/" , serveHome)
	r.HandleFunc("/courses" , getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}" , getOneCourse).Methods("GET")
	r.HandleFunc("/course" , createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}" , updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}" , deleteOneCourse).Methods("DELETE")

	// Listen to a Port
	// Fatal is equivalent to Print() followed by a call to os.Exit(1)
	// ListenAndServe listens on specified TCP network address & then calls Serve with handler to handle requests.
	// The handler is typically nil, in which case the DefaultServeMux is used.
	log.Fatal(http.ListenAndServe(":4000" , r))
}
