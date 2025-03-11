package coinbasego

type ClientCredentials struct {
	KeyName   string
	KeySecret string
}

type Client struct {
	credentials ClientCredentials
}

func NewClient(credentials ClientCredentials) *Client {
	return &Client{
		credentials: credentials,
	}
}
