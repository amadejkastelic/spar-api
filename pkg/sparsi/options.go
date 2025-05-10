package sparsi

import (
	"net/http"
	"net/url"
	"time"
)

type option func(*client)

func WithBaseURL(baseURL string) option {
	return func(c *client) {
		if u, err := url.Parse(baseURL); err == nil {
			c.baseURL = u
		}
	}
}

func WithHttpClient(httpClient *http.Client) option {
	return func(c *client) {
		c.httpClient = httpClient
	}
}

func WithTimeout(timeout time.Duration) option {
	return func(c *client) {
		c.httpClient.Timeout = timeout
	}
}

func WithUserAgent(userAgent string) option {
	return func(c *client) {
		c.userAgent = userAgent
	}
}
