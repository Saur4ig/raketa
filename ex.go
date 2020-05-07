package raketa

import (
	"net/http"

	"github.com/Saur4ig/raketa/rest"
)

func main() {
	httpClient := &http.Client{}
	// Create rest client
	// check err in real project
	restClient, _ := rest.NewClient(
		"127.0.0.1",
		"3000",
		false,
		"testRoom",
		"MyName",
		httpClient,
		false,
	)

	// Login an existing user
	err := restClient.Login(rest.Credentials{Login: "login", Password: "pass"})
	if err != nil {
		panic(err)
	}

	// Send a message
	// check err in real project
	_ = restClient.SendMessage(rest.Message{MessageText: "hello"})
}
