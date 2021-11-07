package dto

type CustomerResponse struct {
	Id        string `json:"customer_id"`
	Name      string `json:"full_name"`
	City      string `json:"city"`
	Zipcode   string `json:"zipcode"`
	BirthDate string `json:"birth_date"`
	Status    string `json:"status"`
}
