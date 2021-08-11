package pkgo

import (
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

// Session is the PluralKit API session, including a token
type Session struct {
	// BaseURL is the API's base url.
	// This is set to the global variables BaseURL + Version when the session is initialized.
	BaseURL string

	Client *http.Client

	token  string
	system *System

	rate *rate.Limiter

	// Timeout is the maximum time this Session will wait for requests.
	Timeout time.Duration
}

// New returns a session with the given token, or no token if the string is empty.
func New(token string) *Session {
	return &Session{
		BaseURL: BaseURL + Version,
		Client:  &http.Client{},
		token:   token,
		rate:    rate.NewLimiter(2, 2),
		Timeout: 10 * time.Second,
	}
}
