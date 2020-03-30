package rest

import (
	"net/http"
	"time"
)

type Client struct {
	Protocol string
	Host     string
	Port     string

	https      bool
	auth       *authInfo
	httpClient *http.Client
}

type authInfo struct {
	token string
	id    string
}

func NewClient(host, port string, tls bool, c *http.Client) *Client {
	var protocol string

	if tls {
		protocol = "https"
	} else {
		protocol = "http"
	}

	if c == nil {
		c = &http.Client{
			Timeout: 5 * time.Second,
		}
	}

	return &Client{Host: host, Port: port, Protocol: protocol, httpClient: c}
}
