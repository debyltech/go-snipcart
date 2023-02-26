package snipcart

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	helper "github.com/debyltech/go-helpers"
	"github.com/skip2/go-qrcode"
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
	Quantity         int     `json:"quantity"`
	TotalWeight      float64 `json:"totalWeight,omitempty"`
	CustomFieldsJSON string  `json:"customFieldsJson"`
	Length           float64 `json:"length,omitempty"`
	Width            float64 `json:"width,omitempty"`
	Height           float64 `json:"height,omitempty"`
	Weight           float64 `json:"weight,omitempty"`
	Shippable        bool    `json:"shippable,omitempty"`
}

type SnipcartOrder struct {
	Token          string         `json:"token"`
	Invoice        string         `json:"invoiceNumber"`
	Status         string         `json:"status"`
	TotalWeight    float64        `json:"totalWeight"`
	Email          string         `json:"email"`
	Name           string         `json:"shippingAddressName"`
	Address1       string         `json:"shippingAddressAddress1"`
	Address2       string         `json:"shippingAddressAddress2"`
	City           string         `json:"shippingAddressCity"`
	Province       string         `json:"shippingAddressProvince"`
	Country        string         `json:"shippingAddressCountry"`
	PostalCode     string         `json:"shippingAddressPostalCode"`
	Phone          string         `json:"shippingAddressPhone"`
	TrackingNumber string         `json:"trackingNumber"`
	TrackingUrl    string         `json:"trackingUrl"`
	ShippingCost   string         `json:"shippingFees"`
	Items          []SnipcartItem `json:"items"`
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

func (s *SnipcartProvider) GetOrder(token string) (*SnipcartOrder, error) {
	response, err := helper.Get(orderUri+"/"+token, "Basic", s.AuthBase64, nil)
	if err != nil {
		return nil, err
	}
	if response.Status != "200 OK" {
		return nil, fmt.Errorf("unexpected response received: %s", response.Status)
	}

	defer response.Body.Close()

	var order SnipcartOrder
	err = json.NewDecoder(response.Body).Decode(&order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (s *SnipcartProvider) GetOrders(queries map[string]string) (*SnipcartOrders, error) {
	response, err := helper.Get(orderUri, "Basic", s.AuthBase64, queries)
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

func (s *SnipcartProvider) GetOrdersByStatus(status OrderStatus) (*SnipcartOrders, error) {
	if status == "" {
		return nil, errors.New("status is not set")
	}

	return s.GetOrders(map[string]string{"status": string(status)})
}

func (o *SnipcartOrder) TokenPNGBase64() (string, error) {
	img, err := qrcode.Encode("order:"+o.Token, qrcode.Medium, 128)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(img), nil
}
