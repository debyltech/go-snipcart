package snipcart

type SnipcartShippingAddress struct {
	FullName    string `json:"fullName"`
	FirstName   string `json:"firstName"`
	Name        string `json:"name"`
	Company     string `json:"company"`
	Address1    string `json:"address1"`
	Address2    string `json:"address2"`
	FullAddress string `json:"fullAddress"`
	City        string `json:"city"`
	Country     string `json:"country"`
	PostalCode  string `json:"postalCode"`
	Province    string `json:"province"`
	Phone       string `json:"phone"`
	VatNumber   string `json:"vatNumber,omitempty"`
}
