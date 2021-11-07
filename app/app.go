package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jwandekoken/golang_rest-server/domain"
	"github.com/jwandekoken/golang_rest-server/logger"
	"github.com/jwandekoken/golang_rest-server/service"
	"github.com/spf13/viper"
)

func setupAndLoadEnv() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	if viper.Get("APP.HOST") == "" ||
		viper.Get("APP.PORT") == "" ||
		viper.Get("DB.HOST") == "" ||
		viper.Get("DB.PORT") == "" ||
		viper.Get("DB.USERNAME") == "" ||
		viper.Get("DB.PASSWORD") == "" ||
		viper.Get("DB.NAME") == "" {
		log.Fatal("Environment variables not defined...")
	}
}

func Start() {
	setupAndLoadEnv()

	router := mux.NewRouter()

	// ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	address := viper.Get("APP.HOST")
	port := viper.Get("APP.PORT")

	// log.Println("Server listening on: ", host)
	logger.Info("Starting the application")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", address, port), router))
}
