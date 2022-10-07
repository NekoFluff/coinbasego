package coinbasego

type Account struct {
	ID             string  `json:"id" binding:"required"`
	Currency       string  `json:"currency" binding:"required"`
	Balance        float64 `json:"balance,string" binding:"required"`
	Available      float64 `json:"available,string" binding:"required"`
	Hold           string  `json:"hold" binding:"required"`
	ProfileId      string  `json:"profile_id" binding:"required"`
	TradingEnabled bool    `json:"trading_enabled" binding:"required"`
}

type AccountsParams struct {
	PaginationParams
}

func (client *Client) Accounts() ([]Account, error) {
	var accounts []Account

	req := Request{
		Method:  "GET",
		PathURL: "/accounts",
		Body:    nil,
	}

	if err := client.sendRequest(req, &accounts, nil); err != nil {
		return nil, err
	}

	return accounts, nil
}

func (client *Client) Account(id string) (*Account, error) {
	var account Account
	req := Request{
		Method:  "GET",
		PathURL: "/accounts/" + id,
		Body:    nil,
	}

	if err := client.sendRequest(req, &account, nil); err != nil {
		return nil, err
	}

	return &account, nil
}
