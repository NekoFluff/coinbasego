package coinbasego

import (
	"net/http"
)

func (client *Client) addAuth(req *http.Request, r Request) error {
	jwt, err := client.JWT(r)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+jwt)

	return nil
}
