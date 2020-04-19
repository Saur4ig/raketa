package rest

import "strings"

// get base url
func (c *Client) getURL() string {
	var sb strings.Builder
	if !c.https {
		sb.WriteString("http://")
	} else {
		sb.WriteString("https://")
	}
	sb.WriteString(c.Host)
	if c.Port != "" {
		sb.WriteString(":")
		sb.WriteString(c.Port)
	}
	return sb.String()
}

// get login url for login action `v1`
func (c *Client) GetLoginURL() string {
	return c.getURL() + "/api/v1/login"
}

// get send message url `v1`
func (c *Client) GetSendMessageURL() string {
	return c.getURL() + "/api/v1/chat.sendMessage"
}
