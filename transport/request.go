package transport

import (
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

func SendRequest(client *http.Client, req *http.Request) (ByteResponse, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("request failed, status %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
