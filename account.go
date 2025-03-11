package coinbasego

type AccountsResponse struct {
	PaginationResponse
	Accounts []Account `json:"accounts"`
}

type AccountResponse struct {
	Account Account `json:"account"`
}

type Account struct {
	UUID             string `json:"uuid" binding:"required"`
	Name             string `json:"name" binding:"required"`
	Currency         string `json:"currency" binding:"required"`
	AvailableBalance struct {
		Value    string `json:"value" binding:"required"`
		Currency string `json:"currency" binding:"required"`
	} `json:"available_balance" binding:"required"`
	Default   bool   `json:"default"`
	Active    bool   `json:"active"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
	Type      string `json:"type"`
	Ready     bool   `json:"ready"`
	Hold      struct {
		Value    string `json:"value" binding:"required"`
		Currency string `json:"currency" binding:"required"`
	}
	RetailPortfolioID string `json:"retail_portfolio_id"`
	Platform          string `json:"platform"`
}

type AccountsParams struct {
	PaginationParams
}

func (client *Client) Accounts() ([]Account, error) {
	var resp AccountsResponse

	req := Request{
		Method:  "GET",
		PathURL: "/api/v3/brokerage/accounts",
		Body:    nil,
	}

	if err := client.sendRequest(req, &resp); err != nil {
		return nil, err
	}

	return resp.Accounts, nil
}

func (client *Client) Account(uuid string) (*Account, error) {
	var resp AccountResponse

	req := Request{
		Method:  "GET",
		PathURL: "/api/v3/brokerage/accounts" + uuid,
		Body:    nil,
	}

	if err := client.sendRequest(req, &resp); err != nil {
		return nil, err
	}

	return &resp.Account, nil
}
