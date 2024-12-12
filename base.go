package exaroton

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type baseClient struct {
	httpc *http.Client
	rb    *requestsBuilder
}

func (c *baseClient) do(req *http.Request) (*http.Response, error) {
	resp, err := c.httpc.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	return resp, nil
}

func decodeRawResponse[T any](resp *http.Response) (*RawResponse[T], error) {
	var dst RawResponse[T]

	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&dst); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if !dst.Success {
		cause := "(no error message provided from API)"
		if dst.Error != nil {
			cause = *dst.Error
		}
		return &dst, fmt.Errorf("API error: [%d] %s", resp.StatusCode, cause)
	}
	return &dst, nil
}
