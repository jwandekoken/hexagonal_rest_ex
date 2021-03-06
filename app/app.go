package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/jwandekoken/hexagonal_rest_ex/domain"
	"github.com/jwandekoken/hexagonal_rest_ex/logger"
	"github.com/jwandekoken/hexagonal_rest_ex/service"
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
	dbClient := getDbClient()

	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)

	ch := CustomerHandlers{service: service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandler{service: service.NewAccountService(accountRepositoryDb)}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).Methods(http.MethodPost)

	address := viper.Get("APP.HOST")
	port := viper.Get("APP.PORT")

	// log.Println("Server listening on: ", host)
	logger.Info("Starting the application")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", address, port), router))
}

func getDbClient() *sqlx.DB {
	dbAddr := viper.Get("DB.HOST")
	dbPort := viper.Get("DB.PORT")
	dbUser := viper.Get("DB.USERNAME")
	dbPwd := viper.Get("DB.PASSWORD")
	dbName := viper.Get("DB.NAME")

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbAddr, dbPort, dbUser, dbPwd, dbName)

	client, err := sqlx.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
