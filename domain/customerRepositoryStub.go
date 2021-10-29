package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Julio", City: "Vitoria", Zipcode: "123123", BirthDate: "30/01/1991", Status: "1"},
		{Id: "1002", Name: "Cesar", City: "SP", Zipcode: "123123", BirthDate: "30/01/1990", Status: "2"},
	}
	return CustomerRepositoryStub{customers}
}
