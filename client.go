package coinbasego

type ClientCredentials struct {
	ApiKey     string
	SecretKey  string
	Passphrase string
}

type Client struct {
	credentials ClientCredentials
}

func NewClient(credentials ClientCredentials) *Client {
	return &Client{
		credentials: credentials,
	}
}
