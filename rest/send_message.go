package rest

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
	"raketa/transport"
	"raketa/types"
)

func (c *Client) Send(payload types.Payload) error {
	if payload.Message.RoomID == "" {
		return errors.New("room id is required")
	}
	if payload.Message.Alias == "" {
		return errors.New("alias is required")
	}

	jsonMessage, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.GetSendMessageURL(), bytes.NewBuffer(jsonMessage))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Token", c.auth.token)
	req.Header.Set("X-User_Id", c.auth.id)

	_, err = transport.SendRequest(c.httpClient, req)
	if err != nil {
		return err
	}

	return nil
}
