package rest

import (
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	type args struct {
		host   string
		port   string
		tls    bool
		roomID string
		alias  string
		c      *http.Client
		debug  bool
	}
	tests := []struct {
		name    string
		args    args
		want    *Client
		wantErr bool
	}{
		{
			name: "Example 1",
			args: args{
				host:   "test.test",
				port:   "007",
				tls:    false,
				roomID: "test_room",
				alias:  "my",
				c:      nil,
				debug:  false,
			},
			want: &Client{
				Host:      "test.test",
				Port:      "007",
				roomID:    "test_room",
				userAlias: "my",
				debug:     false,
				https:     false,
				auth:      nil,
				httpClient: &http.Client{
					Timeout: 5 * time.Second,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClient(tt.args.host, tt.args.port, tt.args.tls, tt.args.roomID, tt.args.alias, tt.args.c, tt.args.debug)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_IsDebug(t *testing.T) {
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
		want   bool
	}{
		{
			name: "Example 1",
			fields: fields{
				Host:       "",
				Port:       "",
				roomID:     "",
				userAlias:  "",
				debug:      false,
				https:      false,
				auth:       nil,
				httpClient: nil,
			},
			want: false,
		},
		{
			name: "Example 2",
			fields: fields{
				Host:       "",
				Port:       "",
				roomID:     "",
				userAlias:  "",
				debug:      true,
				https:      false,
				auth:       nil,
				httpClient: nil,
			},
			want: true,
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
			if got := c.IsDebug(); got != tt.want {
				t.Errorf("IsDebug() = %v, want %v", got, tt.want)
			}
		})
	}
}
