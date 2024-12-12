package exaroton

import (
	"context"
	"encoding/json"

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
func (c *configClient) GetConfig(ctx context.Context) (*model.ConfigFile, *RawResponse[model.ConfigFile], error) {
	resp, err := c.do(c.rb.buildConfigGetRequest(
		ctx,
		c.serverID,
		c.path,
	))
	if err != nil {
		return nil, nil, err
	}

	rawResp, err := decodeRawResponse[model.ConfigFile](resp)
	if err != nil {
		return nil, rawResp, err
	}
	return rawResp.Data, rawResp, nil
}

// SetConfig implements ConfigAPI.
func (c *configClient) SetConfig(ctx context.Context, config json.RawMessage) (*model.ConfigFile, *RawResponse[model.ConfigFile], error) {
	resp, err := c.do(c.rb.buildConfigPostRequest(
		ctx,
		c.serverID,
		c.path,
		config,
	))
	if err != nil {
		return nil, nil, err
	}

	rawResp, err := decodeRawResponse[model.ConfigFile](resp)
	if err != nil {
		return nil, rawResp, err
	}
	return rawResp.Data, rawResp, nil
}
