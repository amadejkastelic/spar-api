package sparsi

import (
	"net/http"
	"net/url"
	"time"
)

type option func(*client)

// WithBaseURL sets the base URL for the Sparsi client.
func WithBaseURL(baseURL string) option {
	return func(c *client) {
		if u, err := url.Parse(baseURL); err == nil {
			c.baseURL = u
		}
	}
}

// WithHttpClient sets the HTTP client for the Sparsi client.
func WithHttpClient(httpClient *http.Client) option {
	return func(c *client) {
		c.httpClient = httpClient
	}
}

// WithTimeout sets the timeout for the HTTP client.
// If not set, the default timeout is 10 seconds.
func WithTimeout(timeout time.Duration) option {
	return func(c *client) {
		c.httpClient.Timeout = timeout
	}
}

// WithUserAgent sets the User-Agent header for the HTTP client.
// If not set, the default User-Agent is "spar-client/1.0".
func WithUserAgent(userAgent string) option {
	return func(c *client) {
		c.userAgent = userAgent
	}
}
