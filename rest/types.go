package rest

// message for creation
type Message struct {
	MessageText string
	Emoji       string
	Avatar      string
	Attachments []Attachment

	alias string
}

type Attachment struct {
	Message   string
	Color     Color
	Title     string
	TitleLink string
}

// sets new alias to message, no whitespaces allowed
func (m *Message) SetAlias(alias string) {
	if alias != "" && alias != " " {
		m.alias = alias
	}
}

// is new alias was set
func (m *Message) IsAliasPresent() bool {
	return m.alias != "" && m.alias != " "
}

// get set message alias
func (m *Message) GetAlias() string {
	return m.alias
}
