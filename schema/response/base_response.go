package response

import "time"

// Base is
type Base struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Timestamp  time.Time   `json:"timestamp"`
	Data       interface{} `json:"data"`
}

// Default for
type Default struct {
	Default interface{} `json:"default,omitempty"`
}
