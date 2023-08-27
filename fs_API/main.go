package main

import (
	"fmt"
	"log"
	"net/http"

	routes "github.com/Sam44323/fs_API/routers"
)

func main() {
	log.Fatal(http.ListenAndServe(":3000", routes.Router()))
	fmt.Printf("Listening on PORT 4000... \n")
}
