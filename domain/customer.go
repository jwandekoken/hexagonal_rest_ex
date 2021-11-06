package domain

import "github.com/jwandekoken/golang_rest-server/errs"

type Customer struct {
	Id        string `db:"customer_id"`
	Name      string
	City      string
	Zipcode   string
	BirthDate string `db:"date_of_birth"`
	Status    string
}

type CustomerRepository interface {
	// status = "1" || "2" || ""
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(id string) (*Customer, *errs.AppError)
}
