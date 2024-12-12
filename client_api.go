package exaroton

import (
	"context"

	"pkg.icikowski.pl/exaroton/model"
)

type apiClient struct {
	*baseClient
}

func newAPIClient(base *baseClient) API {
	return &apiClient{
		baseClient: base,
	}
}

// GetAccount implements API.
func (c *apiClient) GetAccount(ctx context.Context) (*model.AccountInfo, *RawResponse[model.AccountInfo], error) {
	resp, err := c.do(c.rb.buildAccountGetRequest(ctx))
	if err != nil {
		return nil, nil, err
	}

	rawResp, err := decodeRawResponse[model.AccountInfo](resp)
	if err != nil {
		return nil, rawResp, err
	}
	return rawResp.Data, rawResp, nil
}

// GetServers implements API.
func (c *apiClient) GetServers(ctx context.Context) (model.Servers, *RawResponse[model.Servers], error) {
	resp, err := c.do(c.rb.buildServersGetRequest(ctx))
	if err != nil {
		return nil, nil, err
	}

	rawResp, err := decodeRawResponse[model.Servers](resp)
	if err != nil {
		return nil, rawResp, err
	}
	return *rawResp.Data, rawResp, nil
}

// Server implements API.
func (c *apiClient) Server(id string) ServerAPI {
	return newServerClient(c.baseClient, id)
}

// GetCreditPools implements API.
func (c *apiClient) GetCreditPools(ctx context.Context) (model.CreditPools, *RawResponse[model.CreditPools], error) {
	resp, err := c.do(c.rb.buildCreditPoolsGetRequest(ctx))
	if err != nil {
		return nil, nil, err
	}

	rawResp, err := decodeRawResponse[model.CreditPools](resp)
	if err != nil {
		return nil, rawResp, err
	}
	return *rawResp.Data, rawResp, nil
}

// CreditPool implements API.
func (c *apiClient) CreditPool(id string) CreditPoolAPI {
	return newCreditPoolClient(c.baseClient, id)
}
