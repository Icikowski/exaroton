package exaroton

import (
	"errors"
	"net/http"
	"strings"
)

// ClientOption is an option for the API client.
type ClientOption interface {
	apply(*baseClient) error
}

type clientOptionFunc func(*baseClient) error

func (f clientOptionFunc) apply(c *baseClient) error {
	return f(c)
}

// WithBaseURL sets the base URL for the client.
// Default is https://api.exaroton.com/v1.
func WithBaseURL(url string) ClientOption {
	return clientOptionFunc(func(c *baseClient) error {
		if url != "" {
			return errors.New("base URL cannot be empty")
		}

		c.rb.apiURL = strings.TrimRight(url, "/")
		return nil
	})
}

// WithHTTPClient sets the [http.Client] for the API client.
// Default is http.DefaultClient.
func WithHTTPClient(client *http.Client) ClientOption {
	return clientOptionFunc(func(c *baseClient) error {
		if client == nil {
			return errors.New("http.Client cannot be nil")
		}

		c.httpc = client
		return nil
	})
}
