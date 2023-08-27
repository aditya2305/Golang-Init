package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` // will be named as coursename
	Platform string `json:"platform"`
	Price    int
	Password string   `json:"-"`              // will not be included in the JSON
	Tags     []string `json:"tags,omitempty"` // will be included in the JSON only if tags is not empty
}

func main() {
	fmt.Println("Initializing the JSON accessor!")
	EncodeJson()
	DecodeJson()
}

// encoding any data into JSON
func EncodeJson() {

	courses := []course{
		{
			"Go",
			"Golang",
			1000,
			"123",
			[]string{"Go", "Golang", "Programming"},
		},
		{
			"Py",
			"Python",
			2000,
			"456",
			[]string{"Python", "Programming"},
		},
		{
			"JS",
			"Javascript",
			3000,
			"789",
			nil,
		},
	}

	// packaging this data as JSON
	finalJson, err := json.MarshalIndent(courses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(finalJson))
}

func DecodeJson() {

	// decoding a json byte data to json
	jsonData := []byte(`
		{
                "coursename": "Go",
                "platform": "Golang",
                "Price": 1000,
                "tags": ["Go","Golang","Programming"]
    }
	`)

	checkValid := json.Valid(jsonData)

	if checkValid != true {
		fmt.Println("Invalid JSON")
		return
	}

	var courses course
	err := json.Unmarshal(jsonData, &courses) // unmarshal the json data to the course struct

	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", courses) // %#v is used to print interfaces

	// case where the you want key-value pair and don't have a previously created struct for that

	var niceJsonData map[string]interface{} // interface{} when  you don't know the type of the data

	err = json.Unmarshal(jsonData, &niceJsonData)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Nice JSON data: %#v\n", niceJsonData) // %#v is used to print interfaces
}
