package coinbase

import (
	"math"

	"github.com/shopspring/decimal"
)

type Currency struct {
	ID            string   `json:"id" binding:"required"`
	Name          string   `json:"name" binding:"required"`
	MinSize       float64  `json:"min_size,string" binding:"required"`
	Status        string   `json:"status" binding:"required"`
	Message       string   `json:"message"`
	MaxPrecision  float64  `json:"max_precision,string" binding:"required"`
	ConvertibleTo []string `json:"convertible_to"`
	// Details       string   `json:"details" binding:"required"`
}

func (c *Currency) FixSize(size float64) float64 {
	decimalPlaces := int32(math.Round(math.Log10(1.0 / c.MaxPrecision)))
	fixedSize, _ := decimal.NewFromFloat(size).Round(decimalPlaces).Float64()
	return fixedSize
}

func (client *Client) Currencies() ([]Currency, error) {
	var currencies []Currency

	req := Request{
		Method:  "GET",
		PathURL: "/currencies",
		Body:    nil,
	}

	if err := client.sendRequest(req, &currencies, nil); err != nil {
		return nil, err
	}

	return currencies, nil
}

func (client *Client) Currency(id string) (*Currency, error) {
	var currency Currency
	req := Request{
		Method:  "GET",
		PathURL: "/currencies/" + id,
		Body:    nil,
	}

	if err := client.sendRequest(req, &currency, nil); err != nil {
		return nil, err
	}

	return &currency, nil
}
