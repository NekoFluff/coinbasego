package coinbase

import (
	"encoding/json"
	"fmt"
	"math"

	"github.com/google/go-querystring/query"
	"github.com/shopspring/decimal"
)

type Product struct {
	ID              ProductID `json:"id" binding:"required"`
	QuoteCurrency   string    `json:"quote_currency" binding:"required"`
	QuoteIncrement  float64   `json:"quote_increment,string" binding:"required"`
	BaseIncrement   float64   `json:"base_increment,string" binding:"required"`
	DisplayName     string    `json:"display_name" binding:"required"`
	PostOnly        bool      `json:"post_only" binding:"required"`
	LimitOnly       bool      `json:"limit_only" binding:"required"`
	CancelOnly      bool      `json:"cancel_only" binding:"required"`
	Status          string    `json:"status" binding:"required"`
	StatusMessage   string    `json:"status_message" binding:"required"`
	TradingDisabled bool      `json:"trading_disabled"`
}

type ProductsParams struct {
	PaginationParams
}

func (p *Product) FixPrice(size float64) float64 {
	decimalPlaces := int32(math.Round(math.Log10(1.0 / p.QuoteIncrement)))
	fixedPrice, _ := decimal.NewFromFloat(size).Round(decimalPlaces).Float64()
	return fixedPrice
}

func (client *Client) Products() ([]Product, error) {
	var products []Product

	req := Request{
		Method:  "GET",
		PathURL: "/products",
		Body:    nil,
	}

	if err := client.sendRequest(req, &products, nil); err != nil {
		return nil, err
	}

	return products, nil
}

func (client *Client) Product(id string) (*Product, error) {
	var product Product
	req := Request{
		Method:  "GET",
		PathURL: "/products/" + id,
		Body:    nil,
	}

	if err := client.sendRequest(req, &product, nil); err != nil {
		return nil, err
	}

	return &product, nil
}

type Granularity int

var (
	GranularityMinute         Granularity = 60
	GranularityFiveMinutes    Granularity = 300
	GranularityFifteenMinutes Granularity = 900
	GranularityHour           Granularity = 3600
	GranularitySixHours       Granularity = 21600
	GranularityDay            Granularity = 86400
)

type CandlesParams struct {
	PaginationParams
	Granularity Granularity `url:"granularity,omitempty"`
	Start       int64       `url:"start,omitempty"`
	End         int64       `url:"end,omitempty"`
}

type Candle struct {
	Timestamp  int     `json:"timestamp"`
	PriceLow   float64 `json:"price_low"`
	PriceHigh  float64 `json:"price_high"`
	PriceOpen  float64 `json:"price_open"`
	PriceClose float64 `json:"price_close"`
}

func (c *Candle) UnmarshalJSON(b []byte) error {

	var arr []float64

	if err := json.Unmarshal(b, &arr); err != nil {
		return err
	}

	c.Timestamp = int(arr[0])
	c.PriceLow = arr[1]
	c.PriceHigh = arr[2]
	c.PriceOpen = arr[3]
	c.PriceClose = arr[4]

	return nil
}

func (client *Client) Candles(productID string, p CandlesParams) ([]Candle, *PaginationResponse, error) {
	var candles []Candle
	v, _ := query.Values(p)
	paramStr := v.Encode()

	req := Request{
		Method:  "GET",
		PathURL: fmt.Sprintf("/products/%s/candles?%s", productID, paramStr),
		Body:    nil,
	}
	fmt.Println("Candles - Path URL:", req.PathURL)

	pageResp := &PaginationResponse{}
	if err := client.sendRequest(req, &candles, pageResp); err != nil {
		return nil, nil, err
	}

	return candles, pageResp, nil
}
