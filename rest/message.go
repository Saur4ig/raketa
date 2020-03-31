package rest

import (
	"github.com/Saur4ig/raketa/types"
)

func CreateColoredMessage(color types.Color, message, title string) Message {
	return Message{
		Attachments: []Attachment{{
			Message: message,
			Color:   color,
			Title:   title,
		}},
	}
}
