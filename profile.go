package coinbase

import "time"

type Profile struct {
	ID        string    `json:"id" binding:"required"`
	UserID    string    `json:"user_id" binding:"required"`
	Name      string    `json:"name" binding:"required"`
	Active    bool      `json:"active" binding:"required"`
	IsDefault bool      `json:"is_default" binding:"required"`
	HasMargin bool      `json:"has_margin" binding:"required"`
	CreatedAt time.Time `json:"creatd_at" binding:"required"`
}

func (client *Client) Profiles() ([]Profile, error) {
	var accounts []Profile
	req := Request{
		Method:  "GET",
		PathURL: "/profiles",
		Body:    nil,
	}

	if err := client.sendRequest(req, &accounts, nil); err != nil {
		return nil, err
	}

	return accounts, nil
}
