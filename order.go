package coinbasego

import (
	"encoding/json"
	"time"
)

type OrdersResponse struct {
	PaginationResponse
	Orders []Order `json:"orders"`
}

type OrderResponse struct {
	Order Order `json:"order"`
}

type Order struct {
	OrderID            string `json:"order_id"`
	ProductID          string `json:"product_id"`
	UserID             string `json:"user_id"`
	OrderConfiguration struct {
		MarketmarketIOC struct {
			QuoteSize string `json:"quote_size"`
			BaseSize  string `json:"base_size"`
		} `json:"marketmarket_ioc"`
		SorLimitIOC struct {
			QuoteSize  string `json:"quote_size"`
			BaseSize   string `json:"base_size"`
			LimitPrice string `json:"limit_price"`
		} `json:"sor_limit_ioc"`
		LimitLimitGTC struct {
			QuoteSize  string `json:"quote_size"`
			BaseSize   string `json:"base_size"`
			LimitPrice string `json:"limit_price"`
			PostOnly   bool   `json:"post_only"`
		} `json:"limit_limit_gtc"`
		LimitLimitGTD struct {
			QuoteSize  string    `json:"quote_size"`
			BaseSize   string    `json:"base_size"`
			LimitPrice string    `json:"limit_price"`
			EndTime    time.Time `json:"end_time"`
			PostOnly   bool      `json:"post_only"`
		} `json:"limit_limit_gtd"`
		LimitLimitFOK struct {
			QuoteSize  string `json:"quote_size"`
			BaseSize   string `json:"base_size"`
			LimitPrice string `json:"limit_price"`
		} `json:"limit_limit_fok"`
		// TODO: There are others
	} `json:"order_configuration"`
	Side                  OrderSide   `json:"side"`
	ClientOrderID         string      `json:"client_order_id"`
	Status                OrderStatus `json:"status"`
	TimeInForce           TimeInForce `json:"time_in_force"`
	CreatedTime           time.Time   `json:"created_time"`
	CompletionPercentage  string      `json:"completion_percentage"`
	FilledSize            string      `json:"filled_size"`
	AverageFilledPrice    string      `json:"average_filled_price"`
	NumberOfFills         string      `json:"number_of_fills"`
	FilledValue           string      `json:"filled_value"`
	PendingCancel         bool        `json:"pending_cancel"`
	SizeInQuote           bool        `json:"size_in_quote"`
	TotalFees             string      `json:"total_fees"`
	SizeInclusiveOfFees   bool        `json:"size_inclusive_of_fees"`
	TotalValueAfterFees   string      `json:"total_value_after_fees"`
	TriggerStatus         string      `json:"trigger_status"` // TODO: This is an enum
	OrderType             OrderType   `json:"order_type"`
	RejectReason          string      `json:"reject_reason"` // TODO: This is an enum
	Settled               bool        `json:"settled"`
	ProductType           string      `json:"product_type"` // TODO: This is an enum
	RejectMessage         string      `json:"reject_message"`
	CancelMessage         string      `json:"cancel_message"`
	OrderPlacementSource  string      `json:"order_placement_source"` // TODO: This is an enum
	OutstandingHoldAmount string      `json:"outstanding_hold_amount"`
	IsLiquidation         bool        `json:"is_liquidation"`
	LastFillTime          time.Time   `json:"last_fill_time"`
	// TODO: Edit history
	Leverage string `json:"leverage"`
	// TODO: Attached Order (duplicate info of this Order struct)
}

type OrderStatus string

const (
	OrderStatusOpen         OrderStatus = "OPEN"
	OrderStatusPending      OrderStatus = "PENDING"
	OrderStatusFilled       OrderStatus = "FILLED"
	OrderStatusCancelled    OrderStatus = "CANCELLED"
	OrderStatusExpired      OrderStatus = "EXPIRED"
	OrderStatusFailed       OrderStatus = "FAILED"
	OrderStatusUnknown      OrderStatus = "UNKNOWN_ORDER_STATUS"
	OrderStatusQueued       OrderStatus = "QUEUED"
	OrderStatusCancelQueued OrderStatus = "CANCEL_QUEUED"
)

type OrdersParams struct {
	PaginationParams
	OrderIDs     []string      `url:"order_ids,omitempty"`
	ProductIDs   []ProductID   `url:"product_ids,omitempty"`
	OrderStatus  OrderStatus   `url:"order_status,omitempty"`
	TimeInForces []TimeInForce `url:"time_in_forces,omitempty"`
	OrderTypes   []OrderType   `url:"order_types,omitempty"`
	OrderSide    OrderSide     `url:"order_side,omitempty"`
	StartDate    time.Time     `url:"start_date,omitempty"` // RFC3339 timestamp, e.g. 2006-01-02T15:04:05Z
	EndDate      time.Time     `url:"end_date,omitempty"`   // RFC3339 timestamp, e.g. 2006-01-02T15:04:05Z
}

func (client *Client) Orders(p OrdersParams) ([]Order, error) {
	var resp OrdersResponse

	req := Request{
		Method:  "GET",
		PathURL: "/api/v3/brokerage/orders/historical/batch",
		Params:  p,
		Body:    nil,
	}

	if err := client.sendRequest(req, &resp); err != nil {
		return nil, err
	}

	return resp.Orders, nil
}

func (client *Client) Order(id string) (*Order, error) {
	var resp OrderResponse
	req := Request{
		Method:  "GET",
		PathURL: "/api/v3/brokerage/orders/historical/" + id,
		Body:    nil,
	}

	if err := client.sendRequest(req, &resp); err != nil {
		return nil, err
	}

	return &resp.Order, nil
}

type OrderType string

const (
	OrderTypeUnknown   OrderType = "UNKNOWN_ORDER_TYPE"
	OrderTypeMarket    OrderType = "MARKET"
	OrderTypeLimit     OrderType = "LIMIT"
	OrderTypeStop      OrderType = "STOP"
	OrderTypeStopLimit OrderType = "STOP_LIMIT"
	OrderTypeBracket   OrderType = "BRACKET"
	OrderTypeTWAP      OrderType = "TWAP"
)

type OrderSide string

const (
	OrderSideBuy  OrderSide = "BUY"
	OrderSideSell OrderSide = "SELL"
)

type TimeInForce string

const (
	TimeInForceUnknown            TimeInForce = "UNKNOWN_TIME_IN_FORCE"
	TimeInForceGoodUntilDateTime  TimeInForce = "GOOD_UNTIL_DATE_TIME"
	TimeInForceGoodUntilCancelled TimeInForce = "GOOD_UNTIL_CANCELLED"
	TimeInForceImmediateOrCancel  TimeInForce = "IMMEDIATE_OR_CANCEL"
	TimeInForceFillOrKill         TimeInForce = "FILL_OR_KILL"
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

	if err := client.sendRequest(req, &order); err != nil {
		return nil, err
	}

	return &order, nil
}
