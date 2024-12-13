package exaroton

import (
	"context"

	"pkg.icikowski.pl/exaroton/model"
)

type configClient struct {
	*baseClient

	serverID string
	path     string
}

func newConfigClient(base *baseClient, serverID, path string) ConfigAPI {
	return &configClient{
		baseClient: base,
		serverID:   serverID,
		path:       path,
	}
}

// GetConfig implements ConfigAPI.
func (c *configClient) GetConfig(ctx context.Context) (model.ConfigOptions, *RawResponse[model.ConfigOptions], error) {
	resp, err := c.do(c.rb.buildConfigGetRequest(
		ctx,
		c.serverID,
		c.path,
	))
	if err != nil {
		return nil, nil, err
	}

	rawResp, err := decodeRawResponse[model.ConfigOptions](resp)
	if err != nil {
		return nil, rawResp, err
	}
	return *rawResp.Data, rawResp, nil
}

// SetConfig implements ConfigAPI.
func (c *configClient) SetConfig(ctx context.Context, opts model.ConfigValues) (model.ConfigOptions, *RawResponse[model.ConfigOptions], error) {
	resp, err := c.do(c.rb.buildConfigPostRequest(
		ctx,
		c.serverID,
		c.path,
		opts,
	))
	if err != nil {
		return nil, nil, err
	}

	rawResp, err := decodeRawResponse[model.ConfigOptions](resp)
	if err != nil {
		return nil, rawResp, err
	}
	return *rawResp.Data, rawResp, nil
}
