# raketa
[![Build Status](https://travis-ci.org/Saur4ig/raketa.svg?branch=master)](https://travis-ci.org/Saur4ig/raketa)
[![Coverage Status](https://coveralls.io/repos/github/Saur4ig/raketa/badge.svg?branch=master)](https://coveralls.io/github/Saur4ig/raketa?branch=master)  
Rocket.Chat client for messages and notifications, inspired by `https://github.com/detached/gorocket/`

Base example:
```go
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
```