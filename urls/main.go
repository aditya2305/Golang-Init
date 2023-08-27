package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=knvekjdb"

func main() {
	fmt.Println("Initiating the url handling in golang!")

	// parsing the url
	response, _ := url.Parse(myUrl)

	fmt.Println(response.Scheme)
	fmt.Println(response.Host)
	fmt.Println(response.Path)
	fmt.Println(response.RawQuery)
	fmt.Println(response.Port())

	// creating the url with parts
	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/learn",
		RawQuery: "user=sam",
	}

	parseUrl := partsOfUrl.String()
	fmt.Println(parseUrl)
}
