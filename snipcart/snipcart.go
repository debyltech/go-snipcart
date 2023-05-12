package snipcart

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	helper "github.com/debyltech/go-helpers/json"
	"github.com/skip2/go-qrcode"
)

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

type Client struct {
	SnipcartKey string
	AuthBase64  string
	Limit       int
}

type SnipcartCustomField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type SnipcartItem struct {
	UUID         string                `json:"uniqueId"`
	ID           string                `json:"id"`
	Name         string                `json:"name"`
	Quantity     int                   `json:"quantity"`
	TotalWeight  float64               `json:"totalWeight,omitempty"`
	TotalPrice   float64               `json:"totalPrice,omitempty"`
	CustomFields []SnipcartCustomField `json:"customFields"`
	Length       float64               `json:"length,omitempty"`
	Width        float64               `json:"width,omitempty"`
	Height       float64               `json:"height,omitempty"`
	Weight       float64               `json:"weight,omitempty"`
	Shippable    bool                  `json:"shippable,omitempty"`
}

type SnipcartOrder struct {
	Token            string                  `json:"token"`
	Created          time.Time               `json:"creationDate"`
	Modified         time.Time               `json:"modificationDate"`
	Invoice          string                  `json:"invoiceNumber"`
	Subtotal         float64                 `json:"subtotal,omitempty"`
	Currency         string                  `json:"currency,omitempty"`
	Total            float64                 `json:"grandTotal,omitempty"`
	Status           string                  `json:"status"`
	TotalWeight      float64                 `json:"totalWeight"`
	ShippingAddress  SnipcartShippingAddress `json:"shippingAddress,omitempty"`
	Name             string                  `json:"shippingAddressName,omitempty"`
	Company          string                  `json:"shippingAddressCompanyName,omitempty"`
	Address1         string                  `json:"shippingAddressAddress1,omitempty"`
	Address2         string                  `json:"shippingAddressAddress2,omitempty"`
	City             string                  `json:"shippingAddressCity,omitempty"`
	Province         string                  `json:"shippingAddressProvince,omitempty"`
	Country          string                  `json:"shippingAddressCountry,omitempty"`
	PostalCode       string                  `json:"shippingAddressPostalCode,omitempty"`
	Phone            string                  `json:"shippingAddressPhone,omitempty"`
	Email            string                  `json:"email,omitempty"`
	TrackingNumber   string                  `json:"trackingNumber"`
	TrackingUrl      string                  `json:"trackingUrl"`
	ShippingCost     float64                 `json:"shippingFees"`
	ShippingProvider string                  `json:"shippingProvider,omitempty"`
	ShippingMethod   string                  `json:"shippingMethod,omitempty"`
	ShippingRateId   string                  `json:"shippingRateUserDefinedId,omitempty"`
	Items            []SnipcartItem          `json:"items"`
	Metadata         any                     `json:"metadata"`
}

type SnipcartOrderUpdate struct {
	Status         OrderStatus `json:"status"`
	PaymentStatus  string      `json:"paymentStatus,omitempty"`
	TrackingNumber string      `json:"trackingNumber,omitempty"`
	TrackingUrl    string      `json:"trackingUrl,omitempty"`
	ShippingRateId string      `json:"shippingRateUserDefinedId,omitempty"`
	Metadata       any         `json:"metadata,omitempty"`
}

type SnipcartOrders struct {
	TotalItems int
	Items      []SnipcartOrder
}

type SnipcartTax struct {
	Name             string  `json:"name"`
	Amount           float64 `json:"amount"`
	NumberForInvoice string  `json:"numberForInvoice"`
	Rate             float64 `json:"rate"`
}

type SnipcartNotification struct {
	Type           NotificationType `json:"type"`
	DeliveryMethod string           `json:"deliveryMethod"`
	Message        string           `json:"message,omitempty"`
}

type SnipcartNotificationResponse struct {
	Id             string           `json:"id"`
	Created        time.Time        `json:"creationDate"`
	Type           NotificationType `json:"type"`
	DeliveryMethod string           `json:"deliveryMethod"`
	Body           string           `json:"body"`
	Message        string           `json:"message"`
	Subject        string           `json:"subject"`
	SentOn         time.Time        `json:"sentOn"`
}

type SnipcartProductVariant struct {
	Stock          int   `json:"stock"`
	Variation      []any `json:"variation"`
	AllowBackorder bool  `json:"allowOutOfStockPurchases"`
}

type SnipcartProduct struct {
	Token          string                   `json:"id"`
	Id             string                   `json:"userDefinedId"`
	Name           string                   `json:"name"`
	Stock          int                      `json:"stock"`
	TotalStock     int                      `json:"totalStock"`
	AllowBackorder bool                     `json:"allowOutOfStockPurchases"`
	Variants       []SnipcartProductVariant `json:"variants"`
}

type SnipcartProductsResponse struct {
	Keywords      string            `json:"keywords"`
	UserDefinedId string            `json:"userDefinedId"`
	Archived      bool              `json:"archived"`
	From          time.Time         `json:"from"`
	To            time.Time         `json:"to"`
	OrderBy       string            `json:"orderBy"`
	Paginated     bool              `json:"hasMoreResults"`
	TotalItems    int               `json:"totalItems"`
	Offset        int               `json:"offset"`
	Limit         int               `json:"limit"`
	Sort          []any             `json:"sort"`
	Items         []SnipcartProduct `json:"items"`
}

func NewClient(snipcartApiKey string) *Client {
	return &Client{
		SnipcartKey: snipcartApiKey,
		AuthBase64:  base64.StdEncoding.EncodeToString([]byte(snipcartApiKey + ":")),
	}
}

func (s *Client) GetOrder(token string) (*SnipcartOrder, error) {
	response, err := helper.Get(orderUri+"/"+token, "Basic", s.AuthBase64, nil)
	if err != nil {
		return nil, err
	}
	if response.StatusCode < 200 && response.StatusCode >= 300 {
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

func (s *Client) GetOrders(queries map[string]string) (*SnipcartOrders, error) {
	response, err := helper.Get(orderUri, "Basic", s.AuthBase64, queries)
	if err != nil {
		return nil, err
	}
	if response.StatusCode < 200 && response.StatusCode >= 300 {
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

func (s *Client) GetOrdersByStatus(status OrderStatus) (*SnipcartOrders, error) {
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

func (s *Client) UpdateOrder(token string, orderUpdate *SnipcartOrderUpdate) (*SnipcartOrder, error) {
	response, err := helper.Put(orderUri+"/"+token, "Basic", s.AuthBase64, orderUpdate)
	if err != nil {
		return nil, err
	}
	if response.StatusCode < 200 && response.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected response received: %s", response.Status)
	}

	defer response.Body.Close()

	var responseOrder SnipcartOrder
	err = json.NewDecoder(response.Body).Decode(&responseOrder)
	if err != nil {
		return nil, err
	}

	return &responseOrder, nil
}

func (s *Client) SendNotification(token string, notification *SnipcartNotification) (*SnipcartNotificationResponse, error) {
	response, err := helper.Post(orderUri+"/"+token+"/notifications", "Basic", s.AuthBase64, notification)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var responseNotification SnipcartNotificationResponse
	err = json.NewDecoder(response.Body).Decode(&responseNotification)
	if err != nil {
		return nil, err
	}

	return &responseNotification, nil
}

func (s *Client) ValidateWebhook(token string) error {
	validateRequest, err := http.NewRequest("GET", validationUri+token, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}

	auth := base64.StdEncoding.EncodeToString([]byte(s.SnipcartKey + ":"))
	validateRequest.Header.Set("Authorization", fmt.Sprintf("Basic %s", auth))
	validateRequest.Header.Set("Accept", "application/json")

	validateResponse, err := client.Do(validateRequest)
	if err != nil {
		return fmt.Errorf("error validating webhook: %s", err.Error())
	}

	if validateResponse.StatusCode < 200 || validateResponse.StatusCode >= 300 {
		return fmt.Errorf("non-2XX status code for validating webhook: %d", validateResponse.StatusCode)
	}

	return nil
}

func (s *Client) GetProducts(queries map[string]string) (*SnipcartProductsResponse, error) {
	response, err := helper.Get(productsUri, "Basic", s.AuthBase64, queries)
	if err != nil {
		return nil, err
	}
	if response.StatusCode < 200 && response.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected response received: %s", response.Status)
	}

	defer response.Body.Close()

	var products SnipcartProductsResponse
	err = json.NewDecoder(response.Body).Decode(&products)
	if err != nil {
		return nil, err
	}

	return &products, nil
}

func (s *Client) GetProductById(id string) (*SnipcartProduct, error) {
	response, err := helper.Get(productsUri, "Basic", s.AuthBase64, map[string]string{"userDefinedId": id})
	if err != nil {
		return nil, err
	}
	if response.StatusCode < 200 && response.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected response received: %s", response.Status)
	}

	defer response.Body.Close()

	var products SnipcartProductsResponse
	err = json.NewDecoder(response.Body).Decode(&products)
	if err != nil {
		return nil, err
	}

	if len(products.Items) < 1 {
		return nil, fmt.Errorf("no products with id '%s'", id)
	}

	return &products.Items[0], nil
}
