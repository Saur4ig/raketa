package rest

import (
	"github.com/Saur4ig/raketa/types"
)

// message for creation
type Message struct {
	MessageText string
	Emoji       string
	Avatar      string
	Attachments []types.Attachment

	alias string
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
