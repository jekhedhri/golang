package mocked

import (
	"os"
)

// Client -
type Client struct {
	name    string
	baseURL string
	port    string
}

// NewMockClient -
func NewMockClient() *Client {
	client := &Client{
		name:    "MockedClient",
		baseURL: os.Getenv("ENV_mockBaseUrl"),
		port:    os.Getenv("ENV_mockPort"),
	}
	return client
}
