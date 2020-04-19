package rest

// creates message with left side colored pipe
// off doc https://rocket.chat/docs/developer-guides/rest-api/chat/sendmessage/
func CreateColoredMessage(color Color, message, title string) Message {
	return Message{
		Attachments: []Attachment{{
			Message: message,
			Color:   color,
			Title:   title,
		}},
	}
}
