package domain

type Customer struct {
	Id        string
	Name      string
	City      string
	Zipcode   string
	BirthDate string
	Status    string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
