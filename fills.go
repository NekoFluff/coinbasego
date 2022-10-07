package coinbasego

type Fill struct {
	TradeID   string `json:"trade_id" binding:"required"`
	ProductID string `json:"product_id" binding:"required"`
	OrderID   string `json:"order_id" binding:"required"`
	UserID    string `json:"user_id" binding:"required"`
	ProfileID string `json:"profile_id" binding:"required"`
	Liquidity string `json:"liquidity" binding:"required"`
}

func (client *Client) Fills() ([]Fill, error) {
	var fills []Fill
	req := Request{
		Method:  "GET",
		PathURL: "/fills?profile_id=default&limit=100",
		Body:    nil,
	}

	if err := client.sendRequest(req, &fills, nil); err != nil {
		return nil, err
	}

	return fills, nil
}
