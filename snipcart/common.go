package snipcart

const (
	defaultLimit   = 50
	apiUri         = "https://app.snipcart.com"
	ordersPath     = "/api/orders"
	productsPath   = "/api/products"
	validationPath = "/api/requestvalidation/"
)

var (
	orderUri      = apiUri + ordersPath
	productsUri   = apiUri + productsPath
	validationUri = apiUri + validationPath
)

type Address struct {
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

type Client struct {
	Key        string
	AuthBase64 string
	Limit      int
}

type CustomField struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	Type     string `json:"type,omitempty"`
	Options  string `json:"options,omitempty"`
	Required bool   `json:"required"`
}
