package snipcart

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
)

const (
	defaultLimit = 50
	apiUri       = "https://app.snipcart.com"
	ordersPath   = "/api/orders"
)

var (
	orderUri = apiUri + ordersPath
)

type SnipcartProvider struct {
	SnipcartKey string
	AuthBase64  string
	Limit       int
}

type SnipcartItem struct {
	UUID             string  `json:"uniqueId"`
	ID               string  `json:"id"`
	Name             string  `json:"name"`
	Quantity         float64 `json:"quantity"`
	Weight           string  `json:"weight"`
	TotalWeight      float64 `json:"totalWeight"`
	CustomFieldsJSON string  `json:"customFieldsJson"`
}

type SnipcartOrder struct {
	Token       string         `json:"token"`
	Invoice     string         `json:"invoiceNumber"`
	Status      string         `json:"status"`
	TotalWeight float64        `json:"totalWeight"`
	Email       string         `json:"email"`
	Name        string         `json:"shippingAddressName"`
	Address1    string         `json:"shippingAddressAddress1"`
	Address2    string         `json:"shippingAddressAddress2"`
	City        string         `json:"shippingAddressCity"`
	Province    string         `json:"shippingAddressProvince"`
	Country     string         `json:"shippingAddressCountry"`
	PostalCode  string         `json:"shippingAddressPostalCode"`
	Phone       string         `json:"shippingAddressPhone"`
	Items       []SnipcartItem `json:"items"`
}

type SnipcartOrders struct {
	TotalItems int
	Items      []SnipcartOrder
}

func NewSnipcartProvider(snipcartApiKey string) SnipcartProvider {
	return SnipcartProvider{
		SnipcartKey: snipcartApiKey,
		AuthBase64:  base64.StdEncoding.EncodeToString([]byte(snipcartApiKey + ":")),
	}
}

func (s *SnipcartProvider) GetOrders() (*SnipcartOrders, error) {
	return s.GetOrdersByStatus("")
}

func (s *SnipcartProvider) GetOrdersByStatus(status OrderStatus) (*SnipcartOrders, error) {
	if s.Limit == 0 {
		s.Limit = defaultLimit
	}

	queries := []URLQuery{
		{Key: "limit", Value: strconv.Itoa(s.Limit)},
	}

	if status != "" {
		queries = append(queries, URLQuery{Key: "status", Value: string(status)})
	}

	response, err := JSONGet(orderUri, "Basic", s.AuthBase64, queries)
	if err != nil {
		return nil, err
	}
	if response.Status != "200 OK" {
		return nil, fmt.Errorf("unexpected response received: %s", response.Status)
	}

	defer response.Body.Close()

	var orders SnipcartOrders
	err = json.NewDecoder(response.Body).Decode(&orders)
	if err != nil {
		return nil, err
	}

	return &orders, nil
}
