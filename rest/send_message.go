package rest

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Saur4ig/raketa/transport"
	"github.com/Saur4ig/raketa/types"
)

// sends message to the room
// off doc - https://rocket.chat/docs/developer-guides/rest-api/chat/sendmessage/
func (c *Client) SendMessage(mes Message) error {
	payload := types.Payload{
		Message: types.Message{
			RoomID:  c.roomID,
			Alias:   c.userAlias,
			Message: mes.MessageText,
			Emoji:   mes.Emoji,
			Avatar:  mes.Avatar,
		},
	}

	// add attachments to message
	if len(mes.Attachments) > 0 {
		for _, val := range mes.Attachments {
			payload.Message.Attachments = append(payload.Message.Attachments, types.Attachment{
				Color:     val.Color.String(),
				Text:      val.Message,
				Title:     val.Title,
				TitleLink: val.TitleLink,
			})
		}
	}
	// if new alias present - set it
	if mes.IsAliasPresent() {
		payload.Message.Alias = mes.GetAlias()
	}

	jsonMessage, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	if c.IsDebug() {
		log.Println(string(jsonMessage))
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
