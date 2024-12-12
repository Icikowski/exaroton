package exaroton

import "net/http"

// NewClient creates a new Exaroton API client.
func NewClient(token string, opts ...ClientOption) (API, error) {
	base := &baseClient{
		rb:    newRequestBuilder("https://api.exaroton.com/v1", token),
		httpc: http.DefaultClient,
	}

	for _, opt := range opts {
		if err := opt.apply(base); err != nil {
			return nil, err
		}
	}
	return newAPIClient(base), nil
}
