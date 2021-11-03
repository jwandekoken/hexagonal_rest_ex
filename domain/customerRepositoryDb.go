package domain

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// @TODO: put in a env file
const (
	host     = "localhost"
	port     = 5432
	user     = "wandekoken"
	password = "123456"
	dbname   = "golang_rest"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {
	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	rows, err := d.client.Query(findAllSql)
	if err != nil {
		log.Println("Error while querying customers table " + err.Error())
		return nil, err
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.BirthDate, &c.Status)

		if err != nil {
			log.Println("Error while scanning customers table " + err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, error) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = $1"

	row := d.client.QueryRow(customerSql, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.BirthDate, &c.Status)
	if err != nil {
		log.Println("Error while scanning customer " + err.Error())
		return nil, err
	}
	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	client, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client}
}
