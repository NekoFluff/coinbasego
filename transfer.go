package coinbasego

import (
	"fmt"
	"time"

	"github.com/google/go-querystring/query"
)

type TransfersParams struct {
	PaginationParams
	Type string `url:"type,omitempty"`
}

type Transfer struct {
	ID          string                 `url:"id"`
	Type        string                 `url:"type"`
	CreatedAt   time.Time              `url:"created_at"`
	CompletedAt time.Time              `url:"completed_at"`
	CanceledAt  time.Time              `url:"canceled_at"`
	ProcessedAt time.Time              `url:"processed_at"`
	Amount      string                 `url:"amount"`
	UserNonce   string                 `url:"user_nonce"`
	Details     map[string]interface{} `url:"details"`
}

func (client *Client) Transfers(accountID string, p TransfersParams) ([]Transfer, *PaginationResponse, error) {
	var transfer []Transfer
	v, _ := query.Values(p)
	paramStr := v.Encode()

	req := Request{
		Method:  "GET",
		PathURL: fmt.Sprintf("/accounts/%s/transfers?%s", accountID, paramStr),
		Body:    nil,
	}

	pageResp := &PaginationResponse{}
	if err := client.sendRequest(req, &transfer, pageResp); err != nil {
		return nil, nil, err
	}

	return transfer, pageResp, nil
}
