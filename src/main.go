package main

import (
	"controller"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// main function to boot up everything
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/eth/{address}/balance", controller.GetBalance).Methods("GET")
	router.HandleFunc("/eth/{address}/nonce", controller.GetNonce).Methods("GET")
	router.HandleFunc("/eth/{address}", controller.GetAccount).Methods("GET")
	fmt.Println("Start Server on: localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
