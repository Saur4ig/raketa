package types

type Payload struct {
	Message Message `json:"message"`
}

type Message struct {
	RoomID      string       `json:"rid"`
	Message     string       `json:"msg"`
	Alias       string       `json:"alias,omitempty"`
	Emoji       string       `json:"emoji,omitempty"`
	Avatar      string       `json:"avatar,omitempty"`
	Attachments []Attachment `json:"attachments"`
}

type Field struct {
	Short bool   `json:"short,omitempty"`
	Title string `json:"title"`
	Value string `json:"value"`
}

type Attachment struct {
	Color       string  `json:"color"`
	Text        string  `json:"text"`
	Ts          string  `json:"ts,omitempty"`
	ThumbURL    string  `json:"thumb_url,omitempty"`
	MessageLink string  `json:"message_link,omitempty"`
	Collapsed   bool    `json:"collapsed,omitempty"`
	AuthorName  string  `json:"author_name,omitempty"`
	AuthorLink  string  `json:"author_link,omitempty"`
	AuthorIcon  string  `json:"author_icon,omitempty"`
	Title       string  `json:"title,omitempty"`
	TitleLink   string  `json:"title_link,omitempty"`
	ImageURL    string  `json:"image_url,omitempty"`
	Fields      []Field `json:"fields,omitempty"`
}
