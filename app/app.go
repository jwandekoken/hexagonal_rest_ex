package app

import (
	"log"
	"net/http"
)

func Start() {

	// creating or own multiplexer
	mux := http.NewServeMux()

	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomers)

	host := "localhost:8000"

	log.Println("Server listening on: ", host)
	log.Fatal(http.ListenAndServe(host, mux))
}
