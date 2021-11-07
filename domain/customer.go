package domain

import (
	"github.com/jwandekoken/golang_rest-server/dto"
	"github.com/jwandekoken/golang_rest-server/errs"
)

type Customer struct {
	Id        string `db:"customer_id"`
	Name      string
	City      string
	Zipcode   string
	BirthDate string `db:"date_of_birth"`
	Status    string
}

func (c Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:        c.Id,
		Name:      c.Name,
		City:      c.City,
		Zipcode:   c.Zipcode,
		BirthDate: c.BirthDate,
		Status:    c.statusAsText(),
	}
}

type CustomerRepository interface {
	// status = "1" || "2" || ""
	FindAll(string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
