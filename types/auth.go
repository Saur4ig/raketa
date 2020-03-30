package types

// The base for the most of the json responses
type StatusResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type LogonResponse struct {
	StatusResponse
	Data struct {
		Token  string `json:"token"`
		UserID string `json:"userId"`
	}
}
