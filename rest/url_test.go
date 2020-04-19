package rest

import (
	"net/http"
	"testing"
)

func TestClient_getURL(t *testing.T) {
	type fields struct {
		Host       string
		Port       string
		roomID     string
		userAlias  string
		debug      bool
		https      bool
		auth       *authInfo
		httpClient *http.Client
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "with port",
			fields: fields{
				Host:       "test.test",
				Port:       "0808",
				roomID:     "",
				userAlias:  "",
				debug:      false,
				https:      false,
				auth:       nil,
				httpClient: nil,
			},
			want: "http://test.test:0808",
		},
		{
			name: "without port",
			fields: fields{
				Host:       "test.test",
				Port:       "",
				roomID:     "",
				userAlias:  "",
				debug:      false,
				https:      false,
				auth:       nil,
				httpClient: nil,
			},
			want: "http://test.test",
		},
		{
			name: "https",
			fields: fields{
				Host:       "test.test",
				Port:       "",
				roomID:     "",
				userAlias:  "",
				debug:      false,
				https:      true,
				auth:       nil,
				httpClient: nil,
			},
			want: "https://test.test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Host:       tt.fields.Host,
				Port:       tt.fields.Port,
				roomID:     tt.fields.roomID,
				userAlias:  tt.fields.userAlias,
				debug:      tt.fields.debug,
				https:      tt.fields.https,
				auth:       tt.fields.auth,
				httpClient: tt.fields.httpClient,
			}
			if got := c.getURL(); got != tt.want {
				t.Errorf("getURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetLoginURL(t *testing.T) {
	type fields struct {
		Host       string
		Port       string
		roomID     string
		userAlias  string
		debug      bool
		https      bool
		auth       *authInfo
		httpClient *http.Client
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Example 1",
			fields: fields{
				Host:       "test.test",
				Port:       "007",
				roomID:     "",
				userAlias:  "",
				debug:      false,
				https:      false,
				auth:       nil,
				httpClient: nil,
			},
			want: "http://test.test:007/api/v1/login",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Host:       tt.fields.Host,
				Port:       tt.fields.Port,
				roomID:     tt.fields.roomID,
				userAlias:  tt.fields.userAlias,
				debug:      tt.fields.debug,
				https:      tt.fields.https,
				auth:       tt.fields.auth,
				httpClient: tt.fields.httpClient,
			}
			if got := c.GetLoginURL(); got != tt.want {
				t.Errorf("GetLoginURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetSendMessageURL(t *testing.T) {
	type fields struct {
		Host       string
		Port       string
		roomID     string
		userAlias  string
		debug      bool
		https      bool
		auth       *authInfo
		httpClient *http.Client
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Example 1",
			fields: fields{
				Host:       "test.test",
				Port:       "",
				roomID:     "",
				userAlias:  "",
				debug:      false,
				https:      true,
				auth:       nil,
				httpClient: nil,
			},
			want: "https://test.test/api/v1/chat.sendMessage",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Host:       tt.fields.Host,
				Port:       tt.fields.Port,
				roomID:     tt.fields.roomID,
				userAlias:  tt.fields.userAlias,
				debug:      tt.fields.debug,
				https:      tt.fields.https,
				auth:       tt.fields.auth,
				httpClient: tt.fields.httpClient,
			}
			if got := c.GetSendMessageURL(); got != tt.want {
				t.Errorf("GetSendMessageURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
