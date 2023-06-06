package snipcart

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

type TaxShippingInfo struct {
	Fees   float64 `json:"fees"`
	Method string  `json:"method"`
}

type TaxContent struct {
	Created              time.Time       `json:"creationDate"`
	Modified             time.Time       `json:"modificationDate"`
	Token                string          `json:"token"`
	Email                string          `json:"email"`
	ShipToBillingAddress bool            `json:"shipToBillingAddress"`
	BillingAddress       Address         `json:"billingAddress"`
	ShippingAddress      Address         `json:"shippingAddress"`
	InvoiceNumber        string          `json:"invoiceNumber"`
	ShippingInformation  TaxShippingInfo `json:"shippingInformation"`
	Items                []Item          `json:"items"`
	Discounts            []any           `json:"discounts"`
	CustomFields         []CustomField   `json:"customFields"`
	Plans                []any           `json:"plans"`
	Refunds              []any           `json:"refunds"`
	Taxes                []any           `json:"taxes"`
	Currency             string          `json:"currency"`
	Total                float64         `json:"total"`
	DiscountsTotal       float64         `json:"discountsTotal"`
	ItemsTotal           float64         `json:"itemsTotal"`
	TaxesTotal           float64         `json:"taxesTotal"`
	PlansTotal           float64         `json:"plansTotal"`
	TaxProvider          any             `json:"taxProvider"`
	Metadata             any             `json:"metadata"`
}

type TaxWebhook struct {
	Content TaxContent `json:"content"`
}

type Tax struct {
	Name             string  `json:"name"`
	Amount           float64 `json:"amount"`
	NumberForInvoice string  `json:"numberForInvoice"`
	Rate             float64 `json:"rate"`
}

type TaxResponse struct {
	Taxes []Tax `json:"taxes"`
}

func (s *Client) ValidateWebhook(token string) error {
	validateRequest, err := http.NewRequest("GET", validationUri+token, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}

	auth := base64.StdEncoding.EncodeToString([]byte(s.Key + ":"))
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
