package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r)) // like console.error
}

func greeter() {
	fmt.Println("Hello World")
}

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>This is an example header!</h1>"))
}
