package rest

import (
	"testing"
)

func TestMessage_SetAlias(t *testing.T) {
	type fields struct {
		MessageText string
		Emoji       string
		Avatar      string
		Attachments []Attachment
		alias       string
	}
	tests := []struct {
		name   string
		fields fields
		alias  string
		want   string
	}{
		{
			name: "Example 1",
			fields: fields{
				MessageText: "",
				Emoji:       "",
				Avatar:      "",
				Attachments: nil,
				alias:       "",
			},
			alias: "new",
			want:  "new",
		},
		{
			name: "Example 2",
			fields: fields{
				MessageText: "",
				Emoji:       "",
				Avatar:      "",
				Attachments: nil,
				alias:       "old",
			},
			alias: "new",
			want:  "new",
		},
		{
			name: "Example 3",
			fields: fields{
				MessageText: "",
				Emoji:       "",
				Avatar:      "",
				Attachments: nil,
				alias:       "old",
			},
			alias: "",
			want:  "old",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{
				MessageText: tt.fields.MessageText,
				Emoji:       tt.fields.Emoji,
				Avatar:      tt.fields.Avatar,
				Attachments: tt.fields.Attachments,
				alias:       tt.fields.alias,
			}
			m.SetAlias(tt.alias)
			if m.alias != tt.want {
				t.Errorf("SetAlias() = %v, want %v", m.alias, tt.want)
			}
		})
	}
}

func TestMessage_IsAliasPresent(t *testing.T) {
	type fields struct {
		MessageText string
		Emoji       string
		Avatar      string
		Attachments []Attachment
		alias       string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Example 1",
			fields: fields{
				MessageText: "",
				Emoji:       "",
				Avatar:      "",
				Attachments: nil,
				alias:       "",
			},
			want: false,
		},
		{
			name: "Example 2",
			fields: fields{
				MessageText: "",
				Emoji:       "",
				Avatar:      "",
				Attachments: nil,
				alias:       "test",
			},
			want: true,
		},
		{
			name: "Example 3",
			fields: fields{
				MessageText: "",
				Emoji:       "",
				Avatar:      "",
				Attachments: nil,
				alias:       " ",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{
				MessageText: tt.fields.MessageText,
				Emoji:       tt.fields.Emoji,
				Avatar:      tt.fields.Avatar,
				Attachments: tt.fields.Attachments,
				alias:       tt.fields.alias,
			}
			if got := m.IsAliasPresent(); got != tt.want {
				t.Errorf("IsAliasPresent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_GetAlias(t *testing.T) {
	type fields struct {
		MessageText string
		Emoji       string
		Avatar      string
		Attachments []Attachment
		alias       string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Example 1",
			fields: fields{
				MessageText: "",
				Emoji:       "",
				Avatar:      "",
				Attachments: nil,
				alias:       "test 22",
			},
			want: "test 22",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{
				MessageText: tt.fields.MessageText,
				Emoji:       tt.fields.Emoji,
				Avatar:      tt.fields.Avatar,
				Attachments: tt.fields.Attachments,
				alias:       tt.fields.alias,
			}
			if got := m.GetAlias(); got != tt.want {
				t.Errorf("GetAlias() = %v, want %v", got, tt.want)
			}
		})
	}
}
