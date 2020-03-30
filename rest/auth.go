package rest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"raketa/transport"
	"raketa/types"
)

func (c Client) Login(username, pass string) error {
	creds := map[string]string{"user": username, "password": pass}

	jsonCreds, err := json.Marshal(creds)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.GetLoginURL(), bytes.NewBuffer(jsonCreds))
	if err != nil {
		return err
	}

	resp, err := transport.SendRequest(c.httpClient, req)
	if err != nil {
		return err
	}

	logonR := types.LogonResponse{}
	err = json.Unmarshal(resp, &logonR)
	if err != nil {
		return err
	}
	c.auth = &authInfo{id: logonR.Data.UserID, token: logonR.Data.Token}
	return nil
}
