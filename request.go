package coinbasego

import (
	"bytes"
	"net/http"

	"github.com/google/go-querystring/query"
)

type Request struct {
	Method  string
	PathURL string
	Body    []byte
	Params  interface{}
}

func (client *Client) sendRequest(r Request, dataOut interface{}) error {
	url := client.getBaseURL() + r.PathURL
	if r.Params != nil {
		v, _ := query.Values(r.Params)
		url += "?" + v.Encode()
	}
	req, err := http.NewRequest(r.Method, url, bytes.NewReader(r.Body))
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

	return client.process(req, &dataOut)
}

func (client *Client) getBaseURL() string {
	return "https://" + client.getHost()
}

func (client *Client) getHost() string {
	return "api.coinbase.com"
}
