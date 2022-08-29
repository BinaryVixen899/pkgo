package pkgo

import (
	"net/http"
	"time"

	"github.com/joeshaw/fastly-roundtripper/transport"
	"golang.org/x/time/rate"
)

// Session is the PluralKit API session, including a token
type Session struct {
	// BaseURL is the API's base url.
	// This is set to the global variables BaseURL + Version when the session is initialized.
	BaseURL string

	Client *http.Client
	token  string

	rate *rate.Limiter

	// Timeout is the maximum time this Session will wait for requests.
	Timeout time.Duration

	// RequestOptions are applied to every outgoing request.
	RequestOptions []RequestOption
}

// New returns a session with the given token, or no token if the string is empty.
func New(token string) *Session {
	return NewWithLimiter(token, rate.NewLimiter(2, 2))
}

// NewWithLimiter returns a session with the given token and rate limiter.
func NewWithLimiter(token string, limiter *rate.Limiter) *Session {
	t := transport.New("pluralkit")
	t.AddBackend("pluralkit", "api.pluralkit.me")

	s := &Session{
		BaseURL: BaseURL + Version,
		Client:  &http.Client{Transport: t},
		token:   token,
		rate:    limiter,
		Timeout: 10 * time.Second,
	}

	fn := func(req *http.Request) error {
		req.Header.Set("Authorization", s.token)
		return nil
	}

	s.RequestOptions = append(s.RequestOptions, fn)

	return s
}
