package utiles

import (
	"bytes"
	"net/http"
)

type CustomClient struct {
	Client *http.Client
	JWT    string
}

func NewCustomClient(jwt string) *CustomClient {
	return &CustomClient{
		Client: &http.Client{},
		JWT:    jwt,
	}
}

func (c *CustomClient) SendRequest(method, url string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.JWT)
	req.Header.Set("Content-Type", "application/json")

	return c.Client.Do(req)
}
