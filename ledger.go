package coinbasego

import (
	"fmt"
	"time"

	"github.com/google/go-querystring/query"
)

type LedgerParams struct {
	PaginationParams
	ProductID ProductID `url:"product_id,omitempty"`
	StartDate time.Time `url:"start_date,omitempty"`
	EndDate   time.Time `url:"end_date,omitempty"`
	ProfileID string    `url:"profile_id"`
	AccountID string
}

type Ledger struct {
	ID        string                 `url:"id"`
	Amount    string                 `url:"amount"`
	CreatedAt time.Time              `url:"created_at"`
	Balance   string                 `url:"balance"`
	Type      string                 `url:"type"`
	Details   map[string]interface{} `url:"details"`
}

func (client *Client) Ledger(p LedgerParams) ([]Ledger, *PaginationResponse, error) {
	var ledger []Ledger
	v, _ := query.Values(p)
	paramStr := v.Encode()

	req := Request{
		Method:  "GET",
		PathURL: fmt.Sprintf("/accounts/%s/ledger?%s", p.AccountID, paramStr),
		Body:    nil,
	}

	pageResp := &PaginationResponse{}
	if err := client.sendRequest(req, &ledger, pageResp); err != nil {
		return nil, nil, err
	}

	return ledger, pageResp, nil
}
