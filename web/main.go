package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Initiating the request!")

	resp, err := http.Get(url)

	checkNilError(err)

	fmt.Printf("Response is of type: %T\n", resp)

	defer resp.Body.Close()

	databytes, err := ioutil.ReadAll(resp.Body)
	checkNilError(err)

	content := string(databytes)
	fmt.Println("Text-data inside the file is \n", content)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
