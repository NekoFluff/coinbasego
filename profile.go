package coinbasego

type PortfoliosResponse struct {
	Portfolios []Portfolio
}

type Portfolio struct {
	Name    string `json:"name"`
	UUID    string `json:"uuid"`
	Type    string `json:"type"`
	Deleted bool   `json:"deleted"`
}

func (client *Client) Portfolios() ([]Portfolio, error) {
	var resp PortfoliosResponse

	req := Request{
		Method:  "GET",
		PathURL: "/api/v3/brokerage/portfolios",
		Body:    nil,
	}

	if err := client.sendRequest(req, &resp); err != nil {
		return nil, err
	}

	return resp.Portfolios, nil
}
