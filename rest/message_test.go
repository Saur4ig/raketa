package rest

import (
	"reflect"
	"testing"
)

func TestCreateColoredMessage(t *testing.T) {
	type args struct {
		color   Color
		message string
		title   string
	}
	tests := []struct {
		name string
		args args
		want Message
	}{
		{
			name: "Example 1",
			args: args{
				color:   Red,
				message: "my test message",
				title:   "title 11",
			},
			want: Message{
				Emoji:  "",
				Avatar: "",
				Attachments: []Attachment{{
					Message: "my test title",
					Color:   Red,
					Title:   "title 11",
				}},
				alias: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateColoredMessage(tt.args.color, tt.args.message, tt.args.title); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateColoredMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
