package transport

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func testMock(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("test 1"))
}

func serverMock() *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/test", testMock)

	srv := httptest.NewServer(handler)

	return srv
}

func TestSendRequest(t *testing.T) {
	srv := serverMock()
	defer srv.Close()

	req, err := http.NewRequest(http.MethodGet, srv.URL+"/test", nil)
	if err != nil {
		t.Error(err)
		return
	}

	client := &http.Client{}
	t.Run("test 1", func(t *testing.T) {
		res, err := SendRequest(client, req)
		if err != nil {
			t.Error(err)
			return
		}
		if string(res) != "test 1" {
			t.Errorf("SendRequest() = %v, want %v", string(res), "test 1")
		}
	})
}
