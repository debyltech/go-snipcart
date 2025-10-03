package snipcart

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	helper "github.com/debyltech/go-helpers/json"
	"github.com/skip2/go-qrcode"
)

type Item struct {
	UUID         string        `json:"uniqueId"`
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	Quantity     int           `json:"quantity"`
	TotalWeight  float64       `json:"totalWeight,omitempty"`
	TotalPrice   float64       `json:"totalPrice,omitempty"`
	CustomFields []CustomField `json:"customFields"`
	Length       float64       `json:"length,omitempty"`
	Width        float64       `json:"width,omitempty"`
	Height       float64       `json:"height,omitempty"`
	Weight       float64       `json:"weight,omitempty"`
	Shippable    bool          `json:"shippable,omitempty"`
}

type Discount struct {
	ID                       string   `json:"id"`
	DiscountID               string   `json:"discountId"`
	Name                     string   `json:"name"`
	Code                     string   `json:"code,omitempty"`
	Trigger                  string   `json:"trigger,omitempty"`
	Type                     string   `json:"type,omitempty"`
	Rate                     float64  `json:"rate,omitempty"`
	Amount                   float64  `json:"amount,omitempty"`
	AmountSaved              float64  `json:"amountSaved,omitempty"`
	NormalizedRate           float64  `json:"normalizedRate,omitempty"`
	Combinable               bool     `json:"combinable"`
	HasSavedAmount           bool     `json:"hasSavedAmount"`
	MaxDiscountsPerItem      *int     `json:"maxDiscountsPerItem,omitempty"`
	TotalToReach             *float64 `json:"totalToReach,omitempty"`
	ProductIds               string   `json:"productIds,omitempty"`
	Categories               string   `json:"categories,omitempty"`
	NumberOfUsages           int      `json:"numberOfUsages,omitempty"`
	AppliesOnAllRecurring    bool     `json:"appliesOnAllRecurringOrders"`
	CreationDate             string   `json:"creationDate,omitempty"`
	ModificationDate         string   `json:"modificationDate,omitempty"`
}

type OrderTax struct {
	Name             string  `json:"taxName"`
	Rate             float64 `json:"taxRate"`
	Amount           float64 `json:"amount"`
	NumberForInvoice string  `json:"numberForInvoice"`
}

type OrderNotification struct {
	Id             string           `json:"id"`
	Created        time.Time        `json:"creationDate"`
	Type           NotificationType `json:"type"`
	DeliveryMethod string           `json:"deliveryMethod"`
	Message        string           `json:"message,omitempty"`
	SendDate       time.Time        `json:"sentOn,omitempty"`
	Subject        string           `json:"subject,omitempty"`
}

type OrderNotifications struct {
	TotalNotifications int                 `json:"totalItems"`
	Offset             int                 `json:"offset"`
	Limit              int                 `json:"limit"`
	Notifications      []OrderNotification `json:"items"`
}

type Order struct {
	Token                        string     `json:"token"`
	Created                      time.Time  `json:"creationDate"`
	Modified                     time.Time  `json:"modificationDate"`
	Completed                    time.Time  `json:"completionDate"`
	Invoice                      string     `json:"invoiceNumber"`
	Subtotal                     float64    `json:"subtotal,omitempty"`
	Currency                     string     `json:"currency,omitempty"`
	Total                        float64    `json:"grandTotal,omitempty"`
	TotalTaxable                 float64    `json:"taxableTotal,omitempty"`
	TotalTaxes                   float64    `json:"taxesTotal,omitempty"`
	Status                       string     `json:"status"`
	TotalWeight                  float64    `json:"totalWeight"`
	ShippingAddressSameAsBilling bool       `json:"shippingAddressSameAsBilling,omitempty"`
	ShippingAddress              Address    `json:"shippingAddress,omitempty"`
	Name                         string     `json:"shippingAddressName,omitempty"`
	Company                      string     `json:"shippingAddressCompanyName,omitempty"`
	Address1                     string     `json:"shippingAddressAddress1,omitempty"`
	Address2                     string     `json:"shippingAddressAddress2,omitempty"`
	City                         string     `json:"shippingAddressCity,omitempty"`
	Province                     string     `json:"shippingAddressProvince,omitempty"`
	Country                      string     `json:"shippingAddressCountry,omitempty"`
	PostalCode                   string     `json:"shippingAddressPostalCode,omitempty"`
	Phone                        string     `json:"shippingAddressPhone,omitempty"`
	Email                        string     `json:"email,omitempty"`
	TrackingNumber               string     `json:"trackingNumber"`
	TrackingUrl                  string     `json:"trackingUrl"`
	ShippingCost                 float64    `json:"shippingFees"`
	ShippingProvider             string     `json:"shippingProvider,omitempty"`
	ShippingMethod               string     `json:"shippingMethod,omitempty"`
	ShippingRateId               string     `json:"shippingRateUserDefinedId,omitempty"`
	Discounts                    []Discount `json:"discounts,omitempty"`
	SavedAmount                  float64    `json:"savedAmount,omitempty"`
	TotalRebateRate              float64    `json:"totalRebateRate,omitempty"`
	Items                        []Item     `json:"items"`
	Taxes                        []OrderTax `json:"taxes,omitempty"`
	Metadata                     any        `json:"metadata"`
}

type OrderUpdate struct {
	Status         OrderStatus `json:"status"`
	PaymentStatus  string      `json:"paymentStatus,omitempty"`
	TrackingNumber string      `json:"trackingNumber,omitempty"`
	TrackingUrl    string      `json:"trackingUrl,omitempty"`
	ShippingRateId string      `json:"shippingRateUserDefinedId,omitempty"`
	Metadata       any         `json:"metadata,omitempty"`
}

type Orders struct {
	TotalItems int
	Offest     int
	Limit      int
	Items      []Order
}

type Notification struct {
	Type           NotificationType `json:"type"`
	DeliveryMethod string           `json:"deliveryMethod"`
	Message        string           `json:"message,omitempty"`
}

type NotificationResponse struct {
	Id             string           `json:"id"`
	Created        time.Time        `json:"creationDate"`
	Type           NotificationType `json:"type"`
	DeliveryMethod string           `json:"deliveryMethod"`
	Body           string           `json:"body"`
	Message        string           `json:"message"`
	Subject        string           `json:"subject"`
	SentOn         time.Time        `json:"sentOn"`
}

type ProductVariant struct {
	Stock          int   `json:"stock"`
	Variation      []any `json:"variation"`
	AllowBackorder bool  `json:"allowOutOfStockPurchases"`
}

type ProductCustomField struct {
	Name         string   `json:"name"`
	Placeholder  string   `json:"placeholder"`
	DisplayValue string   `json:"displayValue"`
	Type         string   `json:"type"`
	Options      string   `json:"options"`
	Required     bool     `json:"required"`
	Value        string   `json:"value"`
	Operation    float64  `json:"operation"`
	OptionsArray []string `json:"optionsArray"`
}

type Product struct {
	Token          string               `json:"id"`
	Id             string               `json:"userDefinedId"`
	Name           string               `json:"name"`
	Stock          int                  `json:"stock"`
	TotalStock     int                  `json:"totalStock"`
	AllowBackorder bool                 `json:"allowOutOfStockPurchases"`
	CustomFields   []ProductCustomField `json:"customFields"`
	Variants       []ProductVariant     `json:"variants"`
}

type ProductsResponse struct {
	Keywords      string    `json:"keywords"`
	UserDefinedId string    `json:"userDefinedId"`
	Archived      bool      `json:"archived"`
	From          time.Time `json:"from"`
	To            time.Time `json:"to"`
	OrderBy       string    `json:"orderBy"`
	Paginated     bool      `json:"hasMoreResults"`
	TotalItems    int       `json:"totalItems"`
	Offset        int       `json:"offset"`
	Limit         int       `json:"limit"`
	Sort          []any     `json:"sort"`
	Items         []Product `json:"items"`
}

func NewClient(snipcartApiKey string) *Client {
	return &Client{
		Key:        snipcartApiKey,
		AuthBase64: base64.StdEncoding.EncodeToString([]byte(snipcartApiKey + ":")),
	}
}

func (s *Client) GetOrder(token string) (*Order, error) {
	response, err := helper.Get(orderUri+"/"+token, "Basic", s.AuthBase64, nil)
	if err != nil {
		return nil, err
	}
	if response.StatusCode < 200 && response.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected response received: %s", response.Status)
	}

	defer response.Body.Close()

	var order Order
	err = json.NewDecoder(response.Body).Decode(&order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (s *Client) GetOrderNotifications(token string) (*OrderNotifications, error) {
	response, err := helper.Get(orderUri+"/"+token+"/"+orderNotificationPath, "Basic", s.AuthBase64, nil)
	if err != nil {
		return nil, err
	}
	if response.StatusCode < 200 && response.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected response received: %s", response.Status)
	}

	defer response.Body.Close()

	var notifications OrderNotifications
	err = json.NewDecoder(response.Body).Decode(&notifications)
	if err != nil {
		return nil, err
	}

	return &notifications, nil
}

func (s *Client) GetOrders(queries map[string]string) (*Orders, error) {
	response, err := helper.Get(orderUri, "Basic", s.AuthBase64, queries)
	if err != nil {
		return nil, err
	}
	if response.StatusCode < 200 && response.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected response received: %s", response.Status)
	}

	defer response.Body.Close()

	var orders Orders
	err = json.NewDecoder(response.Body).Decode(&orders)
	if err != nil {
		return nil, err
	}

	return &orders, nil
}

func (s *Client) GetOrdersByStatus(status OrderStatus) (*Orders, error) {
	if status == "" {
		return nil, errors.New("status is not set")
	}

	return s.GetOrders(map[string]string{"status": string(status)})
}

func (o *Order) TokenPNGBase64() (string, error) {
	img, err := qrcode.Encode("order:"+o.Token, qrcode.Medium, 128)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(img), nil
}

func (s *Client) UpdateOrder(token string, orderUpdate *OrderUpdate) (*Order, error) {
	response, err := helper.Put(orderUri+"/"+token, "Basic", s.AuthBase64, orderUpdate)
	if err != nil {
		return nil, err
	}
	if response.StatusCode < 200 && response.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected response received: %s", response.Status)
	}

	defer response.Body.Close()

	var responseOrder Order
	err = json.NewDecoder(response.Body).Decode(&responseOrder)
	if err != nil {
		return nil, err
	}

	return &responseOrder, nil
}

func (s *Client) SendNotification(token string, notification *Notification) (*NotificationResponse, error) {
	response, err := helper.Post(orderUri+"/"+token+"/notifications", "Basic", s.AuthBase64, notification)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var responseNotification NotificationResponse
	err = json.NewDecoder(response.Body).Decode(&responseNotification)
	if err != nil {
		return nil, err
	}

	return &responseNotification, nil
}

func (s *Client) GetProducts(queries map[string]string) (*ProductsResponse, error) {
	response, err := helper.Get(productsUri, "Basic", s.AuthBase64, queries)
	if err != nil {
		return nil, err
	}
	if response.StatusCode < 200 && response.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected response received: %s", response.Status)
	}

	defer response.Body.Close()

	var products ProductsResponse
	err = json.NewDecoder(response.Body).Decode(&products)
	if err != nil {
		return nil, err
	}

	return &products, nil
}

func (s *Client) GetProductById(id string) (*Product, error) {
	response, err := helper.Get(productsUri, "Basic", s.AuthBase64, map[string]string{"userDefinedId": id})
	if err != nil {
		return nil, err
	}
	if response.StatusCode < 200 && response.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected response received: %s", response.Status)
	}

	defer response.Body.Close()

	var products ProductsResponse
	err = json.NewDecoder(response.Body).Decode(&products)
	if err != nil {
		return nil, err
	}

	if len(products.Items) < 1 {
		return nil, fmt.Errorf("no products with id '%s'", id)
	}

	return &products.Items[0], nil
}
