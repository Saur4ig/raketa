package rest

import (
	"errors"
	"net/http"
	"time"
)

// rest client for base needs
type Client struct {
	Host string
	Port string

	// id of room or channel
	roomID string
	// username could be changed for specific message, if user is bot
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

// creates new rest raketa client
func NewClient(host, port string, tls bool, roomID, alias string, c *http.Client, debug bool) (*Client, error) {
	if roomID == "" || roomID == " " {
		return nil, errors.New("room id is required")
	}
	if alias == "" || alias == " " {
		return nil, errors.New("alias is required")
	}

	if c == nil {
		c = &http.Client{
			Timeout: 5 * time.Second,
		}
	}

	return &Client{
		https:      tls,
		Host:       host,
		Port:       port,
		httpClient: c,
		roomID:     roomID,
		userAlias:  alias,
		debug:      debug,
	}, nil
}

// checks is client in debug mod
func (c *Client) IsDebug() bool {
	return c.debug
}
