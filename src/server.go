package main

import (
	"fmt"
	"log"
	"net/http"
	"router"
)

// main function to boot up everything
func main() {
	router := router.InitRouter()
	fmt.Println("Start Server on: localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
