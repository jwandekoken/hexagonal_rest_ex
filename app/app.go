package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jwandekoken/golang_rest-server/domain"
	"github.com/jwandekoken/golang_rest-server/service"
)

func Start() {
	router := mux.NewRouter()

	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	host := "localhost:8000"
	log.Println("Server listening on: ", host)
	log.Fatal(http.ListenAndServe(host, router))
}
