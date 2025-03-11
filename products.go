package coinbasego

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
)

type Product struct {
	ProductID                 string `json:"product_id"`
	Price                     string `json:"price"`
	PricePercentageChange24H  string `json:"price_percentage_change_24h"`
	Volume24H                 string `json:"volume_24h"`
	VolumePercentageChange24H string `json:"volume_percentage_change_24h"`
	BaseIncrement             string `json:"base_increment"`
	QuoteIncrement            string `json:"quote_increment"`
	QuoteMinSize              string `json:"quote_min_size"`
	QuoteMaxSize              string `json:"quote_max_size"`
	BaseMinSize               string `json:"base_min_size"`
	BaseName                  string `json:"base_name"`
	QuoteName                 string `json:"quote_name"`
	Watched                   bool   `json:"watched"`
	IsDisabled                bool   `json:"is_disabled"`
	New                       bool   `json:"new"`
	Status                    string `json:"status"`
	CancelOnly                bool   `json:"cancel_only"`
	LimitOnly                 bool   `json:"limit_only"`
	PostOnly                  bool   `json:"post_only"`
	TradingDisabled           bool   `json:"trading_disabled"`
	AuctionMode               bool   `json:"auction_mode"`
	ProductType               string `json:"product_type"`
	QuoteCurrencyID           string `json:"quote_currency_id"`
	BaseCurrencyID            string `json:"base_currency_id"`
	// TODO: FCM Trading Session Details
}

type ProductsParams struct {
	PaginationParams
}

func (p *Product) FixPrice(size float64) float64 {
	quoteIncrement, _ := strconv.ParseFloat(p.QuoteIncrement, 64)
	decimalPlaces := int32(math.Round(math.Log10(1.0 / quoteIncrement)))
	fixedPrice, _ := decimal.NewFromFloat(size).Round(decimalPlaces).Float64()
	return fixedPrice
}

func (client *Client) Products() ([]Product, error) {
	var resp struct {
		Products    []Product `json:"products"`
		NumProducts int32     `json:"num_products"`
	}

	req := Request{
		Method:  "GET",
		PathURL: "/api/v3/brokerage/products",
		Body:    nil,
	}

	if err := client.sendRequest(req, &resp); err != nil {
		return nil, err
	}

	return resp.Products, nil
}

func (client *Client) Product(id string) (*Product, error) {
	var product Product
	req := Request{
		Method:  "GET",
		PathURL: "/api/v3/brokerage/products/" + id,
		Body:    nil,
	}

	if err := client.sendRequest(req, &product); err != nil {
		return nil, err
	}

	return &product, nil
}

type Granularity string

var (
	GranularityMinute         Granularity = "ONE_MINUTE"
	GranularityFiveMinutes    Granularity = "FIVE_MINUTE"
	GranularityFifteenMinutes Granularity = "FIFTEEN_MINUTE"
	GranularityHour           Granularity = "ONE_HOUR"
	GranularitySixHours       Granularity = "SIX_HOUR"
	GranularityDay            Granularity = "ONE_DAY"
)

type CandlesParams struct {
	Limit       int32       `url:"limit,omitempty"`
	Granularity Granularity `url:"granularity,omitempty"`
	Start       int64       `url:"start,omitempty"`
	End         int64       `url:"end,omitempty"`
}

type CandlesResponse struct {
	Candles []Candle `json:"candles"`
}

type Candle struct {
	Start  time.Time `json:"start"`
	Low    float64   `json:"low"`
	High   float64   `json:"high"`
	Open   float64   `json:"open"`
	Close  float64   `json:"close"`
	Volume string    `json:"volume"`
}

func (c *Candle) UnmarshalJSON(b []byte) error {
	var values struct {
		Start  string `json:"start"`
		Low    string `json:"low"`
		High   string `json:"high"`
		Open   string `json:"open"`
		Close  string `json:"close"`
		Volume string `json:"volume"`
	}

	var err error
	if err = json.Unmarshal(b, &values); err != nil {
		return err
	}

	ts, err := strconv.ParseInt(values.Start, 10, 64)
	if err != nil {
		return err
	}

	c.Start = time.Unix(ts, 0)

	c.Low, err = strconv.ParseFloat(values.Low, 64)
	if err != nil {
		return err
	}

	c.High, err = strconv.ParseFloat(values.High, 64)
	if err != nil {
		return err
	}

	c.Open, err = strconv.ParseFloat(values.Open, 64)
	if err != nil {
		return err
	}

	c.Close, err = strconv.ParseFloat(values.Close, 64)
	if err != nil {
		return err
	}

	c.Volume = values.Volume

	return nil
}

func (client *Client) Candles(productID string, p CandlesParams) ([]Candle, error) {
	var resp CandlesResponse

	req := Request{
		Method:  "GET",
		PathURL: fmt.Sprintf("/api/v3/brokerage/products/%s/candles", productID),
		Params:  p,
		Body:    nil,
	}

	if err := client.sendRequest(req, &resp); err != nil {
		return nil, err
	}

	return resp.Candles, nil
}
