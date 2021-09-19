package securitytrails

const BaseURL = 'https://api.securitytrails.com/v1/'

type Client struct{
	apiKey string
}

func New(apiKey string) *Client{
	return &Client{apiKey: apiKey}
}
