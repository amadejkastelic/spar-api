package sparsi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/amadejkastelic/spar-api/internal/utils"
)

const (
	defaultBaseURL = "https://search-spar.spar-ics.com/fact-finder/rest/v4/"
	defaultTimeout = 10 * time.Second
)

type Client interface {
	Categories() (*CategoriesResponse, error)
	Search(opts SearchRequest) (*SearchResponse, error)
}

type client struct {
	baseURL    *url.URL
	httpClient *http.Client
	userAgent  string
}

// NewClient creates a new Sparsi client with the given options.
func NewClient(opts ...option) *client {
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &client{
		baseURL:    baseURL,
		httpClient: &http.Client{Timeout: defaultTimeout},
		userAgent:  "spar-client/1.0",
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *client) newRequest(
	method, path string,
	query url.Values,
	body any,
) (*http.Request, error) {
	rel := &url.URL{Path: path, RawQuery: query.Encode()}
	u := c.baseURL.ResolveReference(rel)

	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	req, err := http.NewRequest(method, u.String(), bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (c *client) do(req *http.Request, v any) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer utils.CloseQuetly(resp.Body)

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return resp, fmt.Errorf("HTTP request failed with status: %s", resp.Status)
	}

	if v != nil {
		if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
			return resp, err
		}
	}

	return resp, nil
}
