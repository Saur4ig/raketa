package rest

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/Saur4ig/raketa/transport"
	"github.com/Saur4ig/raketa/types"
)

func (c *Client) Send(mes Message) error {
	payload := types.Payload{
		Message: types.Message{
			RoomID:      c.roomID,
			Alias:       c.userAlias,
			Message:     mes.MessageText,
			Emoji:       mes.Emoji,
			Avatar:      mes.Avatar,
			Attachments: mes.Attachments,
		},
	}
	// if new alias present - set it
	if mes.IsAliasPresent() {
		payload.Message.Alias = mes.GetAlias()
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
	req.Header.Set("X-User-Id", c.auth.id)

	_, err = transport.SendRequest(c.httpClient, req)
	if err != nil {
		return err
	}

	return nil
}
