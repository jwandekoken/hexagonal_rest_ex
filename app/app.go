package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	// creating a gorilla/mux
	mux := mux.NewRouter()

	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomers)

	host := "localhost:8000"

	log.Println("Server listening on: ", host)
	log.Fatal(http.ListenAndServe(host, mux))
}
