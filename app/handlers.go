package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/jwandekoken/golang_rest-server/service"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city"  xml:"city"`
	Zipcode string `json:"zip_code"  xml:"zipcode"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{Name: "Julio", City: "Vitoria", Zipcode: "2901231"},
	// 	{Name: "Rob", City: "Vitoria", Zipcode: "2901231"},
	// }

	customers, _ := ch.service.GetAllCustomers()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
