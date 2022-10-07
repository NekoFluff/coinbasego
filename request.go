package coinbasego

import (
	"bytes"
	"net/http"
)

type Request struct {
	Method  string
	PathURL string
	Body    []byte
}

func (client *Client) sendRequest(r Request, dataOut interface{}, pageOut *PaginationResponse) error {
	req, err := http.NewRequest(r.Method, client.getBaseURL()+r.PathURL, bytes.NewReader(r.Body))
	if err != nil {
		return err
	}

	err = client.addAuth(req, r)
	if err != nil {
		return err
	}

	if r.Body != nil {
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Content-Type", "application/json")
	}

	return client.process(req, &dataOut, pageOut)
}

func (client *Client) getBaseURL() string {
	return "https://api.exchange.coinbasego.com"
}
