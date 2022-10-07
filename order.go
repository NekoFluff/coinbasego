package coinbasego

import (
	"encoding/json"
	"time"

	"github.com/google/go-querystring/query"
)

type Order struct {
	ID             string      `json:"id" binding:"required"`
	Price          string      `json:"price"`
	Size           string      `json:"size"`
	ProductID      ProductID   `json:"product_id" binding:"required"`
	ProfileID      string      `json:"profile_id"`
	Side           OrderSide   `json:"side" binding:"required"`
	Funds          string      `json:"funds"`
	SpecifiedFunds string      `json:"specified_funds"`
	Type           OrderType   `json:"type" binding:"required"`
	TimeInForce    string      `json:"time_in_force"`
	ExpireTime     string      `json:"expire_time"`
	PostOnly       bool        `json:"post_only" binding:"required"`
	CreatedAt      string      `json:"created_at" binding:"required"`
	DoneAt         string      `json:"done_at"`
	DoneReason     string      `json:"done_reason"`
	RejectReason   string      `json:"reject_reason"`
	FillFees       string      `json:"fill_fees" binding:"required"`
	FilledSize     string      `json:"filled_size" binding:"required"`
	ExecutedValue  string      `json:"executed_value"`
	Status         OrderStatus `json:"status" binding:"required"`
	Settled        bool        `json:"settled" binding:"required"`
	Stop           string      `json:"stop"`
	StopPrice      string      `json:"stop_price"`
	FundingAmount  string      `json:"funding_amount"`
	ClientOID      string      `json:"client_oid"`
}

type Sorting string

const (
	SortingAsc  Sorting = "asc"
	SortingDesc Sorting = "desc"
)

type OrderStatus string

const (
	OrderStatusOpen     OrderStatus = "open"
	OrderStatusPending  OrderStatus = "pending"
	OrderStatusRejected OrderStatus = "rejected"
	OrderStatusDone     OrderStatus = "done"
	OrderStatusActive   OrderStatus = "active"
	OrderStatusReceived OrderStatus = "received"
	OrderStatusAll      OrderStatus = "all"
)

type OrdersParams struct {
	PaginationParams
	ProfileID string        `url:"profile_id,omitempty"`
	ProductID ProductID     `url:"product_id,omitempty"`
	SortedBy  string        `url:"sortedBy,omitempty"`
	Sorting   Sorting       `url:"sorting,omitempty"`
	StartDate time.Time     `url:"start_date,omitempty"`
	EndDate   time.Time     `url:"end_date,omitempty"`
	Status    []OrderStatus `url:"status" binding:"required"`
}

func (client *Client) Orders(p OrdersParams) ([]Order, *PaginationResponse, error) {
	var orders []Order
	v, _ := query.Values(p)
	paramStr := v.Encode()

	req := Request{
		Method:  "GET",
		PathURL: "/orders?" + paramStr,
		Body:    nil,
	}

	pageResp := &PaginationResponse{}
	if err := client.sendRequest(req, &orders, pageResp); err != nil {
		return nil, nil, err
	}

	return orders, pageResp, nil
}

func (client *Client) Order(id string) (*Order, error) {
	var order Order
	req := Request{
		Method:  "GET",
		PathURL: "/orders/" + id,
		Body:    nil,
	}

	if err := client.sendRequest(req, &order, nil); err != nil {
		return nil, err
	}

	return &order, nil
}

type OrderType string

const (
	OrderTypeLimit  OrderType = "limit"
	OrderTypeMarket OrderType = "market"
	OrderTypeStop   OrderType = "stop"
)

type OrderSide string

const (
	OrderSideBuy  OrderSide = "buy"
	OrderSideSell OrderSide = "sell"
)

type TimeInForce string

const (
	TimeInForceGoodTillCanceled = "GTC"
	TimeInForceGoodTillTime     = "GTT"
)

type CancelAfter string

const (
	CancelAfterMin  = "min"
	CancelAFterHour = "hour"
	CancelAfterDay  = "day"
)

type OrderCreateBody struct {
	ProfileID   string      `json:"profile_id,omitempty"`
	Type        OrderType   `json:"type,omitempty"`
	Side        OrderSide   `json:"side" binding:"required"`
	ProductID   ProductID   `json:"product_id" binding:"required"`
	STP         string      `json:"stp,omitempty"`
	Stop        string      `json:"stop,omitempty"`
	StopPrice   string      `json:"stop_price,omitempty"`
	Price       string      `json:"price,omitempty"`
	Size        string      `json:"size,omitempty"`
	Funds       string      `json:"funds,omitempty"`
	TimeInForce TimeInForce `json:"time_in_force,omitempty"`
	CancelAfter CancelAfter `json:"cancel_after,omitempty"`
	PostOnly    bool        `json:"post_only,omitempty"`
	ClientOID   string      `json:"client_oid,omitempty"`
}

func (client *Client) OrderCreate(body OrderCreateBody) (*Order, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var order Order
	req := Request{
		Method:  "POST",
		PathURL: "/orders",
		Body:    b,
	}

	if err := client.sendRequest(req, &order, nil); err != nil {
		return nil, err
	}

	return &order, nil
}
