package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jwandekoken/golang_rest-server/domain"
	"github.com/jwandekoken/golang_rest-server/logger"
	"github.com/jwandekoken/golang_rest-server/service"
)

func Start() {
	router := mux.NewRouter()

	// ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	host := "localhost:8000"
	// log.Println("Server listening on: ", host)
	logger.Info("Starting the application")
	log.Fatal(http.ListenAndServe(host, router))
}
