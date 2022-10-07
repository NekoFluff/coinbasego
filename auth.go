package coinbase

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strconv"
	"time"
)

func (client *Client) addAuth(req *http.Request, r Request) error {
	cb_access_timestamp := strconv.Itoa(int(time.Now().UTC().Unix()))
	cb_access_sign, err := client.accessSignValue(cb_access_timestamp, r)
	if err != nil {
		return err
	}

	req.Header.Set("CB-ACCESS-TIMESTAMP", cb_access_timestamp)
	req.Header.Set("CB-ACCESS-SIGN", cb_access_sign)
	req.Header.Set("CB-ACCESS-KEY", client.credentials.ApiKey)
	req.Header.Set("CB-ACCESS-PASSPHRASE", client.credentials.Passphrase)

	return nil
}

func (client *Client) accessSignValue(timestamp string, r Request) (string, error) {
	message := timestamp + r.Method + r.PathURL
	if r.Body != nil {
		message += string(r.Body)
	}

	decodedSecret, err := base64.StdEncoding.DecodeString(client.credentials.SecretKey)
	if err != nil {
		return "", err
	}
	h := hmac.New(sha256.New, []byte(decodedSecret))
	h.Write([]byte(message))
	hexMessage := h.Sum(nil)

	return base64.StdEncoding.EncodeToString([]byte(hexMessage)), nil
}
