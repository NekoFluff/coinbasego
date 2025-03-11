package coinbasego

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) process(req *http.Request, dataOut interface{}) error {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("unexpected error response (%d): %s", resp.StatusCode, string(b))
	}

	if dataOut != nil {
		if err = json.Unmarshal(b, &dataOut); err != nil {
			return err
		}
	}

	return nil
}
