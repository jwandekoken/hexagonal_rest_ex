package domain

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "wandekoken"
	password = "123456"
	dbname   = "golang_rest"
)

type CustomerRepositoryDb struct {
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	customers := []Customer{
		{Id: "1001", Name: "Julio", City: "Vitoria", Zipcode: "123123", BirthDate: "30/01/1991", Status: "1"},
		{Id: "1002", Name: "Cesar", City: "SP", Zipcode: "123123", BirthDate: "30/01/1990", Status: "2"},
	}

	return customers, nil
}
