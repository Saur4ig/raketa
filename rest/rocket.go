package rest

import (
	"errors"
	"net/http"
	"time"
)

type Client struct {
	Protocol string
	Host     string
	Port     string

	roomID    string
	userAlias string

	debug      bool
	https      bool
	auth       *authInfo
	httpClient *http.Client
}

type authInfo struct {
	token string
	id    string
}

func NewClient(host, port string, tls bool, roomID, alias string, c *http.Client, debug bool) (*Client, error) {
	if roomID == "" || roomID == " " {
		return nil, errors.New("room id is required")
	}
	if alias == "" || alias == " " {
		return nil, errors.New("alias is required")
	}

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

	return &Client{
		Host:       host,
		Port:       port,
		Protocol:   protocol,
		httpClient: c,
		roomID:     roomID,
		userAlias:  alias,
		debug:      debug,
	}, nil
}

func (c *Client) IsDebug() bool {
	return c.debug
}
