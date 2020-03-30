package types

import "time"

type Payload struct {
	Message Message `json:"message"`
}
type Message struct {
	RoomID      string     `json:"rid"`
	Message     string     `json:"msg"`
	Alias       string     `json:"alias"`
	Emoji       string     `json:"emoji"`
	Avatar      string     `json:"avatar"`
	Attachments Attachment `json:"attachments"`
}

type Field struct {
	Short bool   `json:"short"`
	Title string `json:"title"`
	Value string `json:"value"`
}

type Attachment struct {
	Color       string    `json:"color"`
	Text        string    `json:"text"`
	Ts          time.Time `json:"ts"`
	ThumbURL    string    `json:"thumb_url"`
	MessageLink string    `json:"message_link"`
	Collapsed   bool      `json:"collapsed"`
	AuthorName  string    `json:"author_name"`
	AuthorLink  string    `json:"author_link"`
	AuthorIcon  string    `json:"author_icon"`
	Title       string    `json:"title"`
	TitleLink   string    `json:"title_link"`
	ImageURL    string    `json:"image_url"`
	Fields      []Field   `json:"fields"`
}

func CreateMessage(roomID, alias, message string) *Payload {
	return &Payload{Message: Message{
		RoomID:      roomID,
		Message:     message,
		Alias:       alias,
		Emoji:       "",
		Avatar:      "",
		Attachments: Attachment{},
	}}
}

func (m *Payload) SetAvatar(url string) {
	m.Message.Avatar = url
}
