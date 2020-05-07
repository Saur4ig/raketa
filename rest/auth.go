package rest

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/Saur4ig/raketa/transport"
)

// The base for the most of the json responses
type StatusResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// response for login request
type LogonResponse struct {
	StatusResponse
	Data struct {
		Token  string `json:"authToken"`
		UserID string `json:"userId"`
	}
}

// credentials for login
type Credentials struct {
	Login    string
	Password string
}

// Login a user
// off docs - https://rocket.chat/docs/developer-guides/rest-api/authentication/login/
func (c *Client) Login(cr Credentials) error {
	credentials := map[string]string{"user": cr.Login, "password": cr.Password}

	jsonCredentials, err := json.Marshal(credentials)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.GetLoginURL(), bytes.NewBuffer(jsonCredentials))
	if err != nil {
		return err
	}

	resp, err := transport.SendRequest(c.httpClient, req)
	if err != nil {
		return err
	}

	logonR := LogonResponse{}
	err = json.Unmarshal(resp, &logonR)
	if err != nil {
		return err
	}
	c.auth = &authInfo{id: logonR.Data.UserID, token: logonR.Data.Token}
	return nil
}
