package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Initializing the web_requester!")
	PerformGetRequest()
	PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	url := "http://localhost:3000/"

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("Status_code for the response: ", response.StatusCode)

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println("Content of the response: ", string(content))

	/*
		Alternative method for reading a response
		var responseString string.Builder
		content, _ := ioutil.ReadAll(response.Body)
		byteCount, _ := responseString.Write(content)

		fmt.Println("Content of the response: ", responseString.String()) // printing the string content of response
	*/
}

func PerformPostJsonRequest() {
	url := "http://localhost:3000/post"

	// json fake data using ``

	requestBody := strings.NewReader(`
	{"data": "This is a fake json data"}
	`)

	response, err := http.Post(url, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Content of the response: ", string(content))
}

func PerformPostFormRequest() {
	myUrl := "http://localhost:3000/postForm"

	// creating a form_data
	data := url.Values{}
	data.Add("firstname", "Sam")
	data.Add("lastname", "Henrick")
	data.Add("email", "sam@henrick.com")

	response, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Content of the response: ", string(content))
}
