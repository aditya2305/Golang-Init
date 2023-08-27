/*
Marshal and Unmarshal convert a string into JSON and vice versa. Encoding and decoding convert a stream into JSON and vice versa.
*/

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

// Models for course and author

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middlewares

func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

// controllers

// serve the home_route

func serveHome(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("<h1>Welcome to the Home Page</h1>"))
}

// getting all the course
func getCourses(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Getting all the data")
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(courses)
}

// getting a single course
func getOneCourse(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Getting one course")
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req) // getting the id from the url
	for _, c := range courses {
		if c.CourseId == params["id"] {
			json.NewEncoder(res).Encode(c)
			return
		}
	}
	json.NewEncoder(res).Encode("No such course was found!")
}

// creating a new_course
func createOneCourse(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Creating one course")
	res.Header().Set("Content-Type", "application/json")

	// if the request data is empty
	if req.Body == nil {
		json.NewEncoder(res).Encode("No data was sent!")
		return
	}

	// when the data is {}
	var course Course
	_ = json.NewDecoder(req.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(res).Encode("No data was sent!")
		return
	}

	// when the title is same
	for _, c := range courses {
		if c.CourseName == course.CourseName {
			json.NewEncoder(res).Encode("Course already exists!")
			return
		}
	}

	// generating a new id and converting it to string
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100)) // converting the integer to a string
	courses = append(courses, course)
	json.NewEncoder(res).Encode(course)
	return

	// appending the new course to the courses_slice
}

// updating a course

func updateOneCourse(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Updating one course")
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req) // getting the id from the url
	for index, c := range courses {
		if c.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...) // removing the course with that id
			var course Course
			_ = json.NewDecoder(req.Body).Decode(&course)
			// when the title is same
			for _, c := range courses {
				if c.CourseName == course.CourseName {
					json.NewEncoder(res).Encode("Course already exists!")
					return
				}
			}
			course.CourseId = params["id"]
			courses = append(courses, course) // adding the new/updated course
			json.NewEncoder(res).Encode(course)
			return
		}
	}
	json.NewEncoder(res).Encode("No such course was found!")
}

// deleting a course
func deleteOneCourse(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Updating one course")
	res.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req) // getting the id from the url
	for index, c := range courses {
		if c.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...) // removing the course with that id
			json.NewEncoder(res).Encode(courses)
			return
		}
	}
	json.NewEncoder(res).Encode("No such course was found!")
}

func main() {
	fmt.Println("Initiating the server")

	// creating a new router
	r := mux.NewRouter()

	// seeding the database
	courses = append(courses, Course{CourseId: "1", CourseName: "Go", CoursePrice: 100, Author: &Author{Fullname: "Golang", Website: "golang.org"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "Python", CoursePrice: 200, Author: &Author{Fullname: "Python", Website: "python.org"}})
	courses = append(courses, Course{CourseId: "3", CourseName: "Java", CoursePrice: 300, Author: &Author{Fullname: "Java", Website: "java.org"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/courses", createOneCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", deleteOneCourse).Methods("DELETE")

	// listening to a port
	log.Fatal(http.ListenAndServe(":3000", r)) // equivalent to console.error
}
