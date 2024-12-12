package exaroton

import (
	"context"

	"pkg.icikowski.pl/exaroton/model"
)

type creditPoolClient struct {
	*baseClient

	poolID string
}

func newCreditPoolClient(base *baseClient, poolID string) CreditPoolAPI {
	return &creditPoolClient{
		baseClient: base,
		poolID:     poolID,
	}
}

// GetCreditPool implements CreditPoolAPI.
func (c *creditPoolClient) GetCreditPool(ctx context.Context) (*model.CreditPool, *RawResponse[model.CreditPool], error) {
	resp, err := c.do(c.rb.buildCreditPoolGetRequest(
		ctx,
		c.poolID,
	))
	if err != nil {
		return nil, nil, err
	}

	rawResp, err := decodeRawResponse[model.CreditPool](resp)
	if err != nil {
		return nil, rawResp, err
	}
	return rawResp.Data, rawResp, nil
}

// GetServers implements CreditPoolAPI.
func (c *creditPoolClient) GetServers(ctx context.Context) (model.Servers, *RawResponse[model.Servers], error) {
	resp, err := c.do(c.rb.buildCreditPoolServersGetRequest(
		ctx,
		c.poolID,
	))
	if err != nil {
		return nil, nil, err
	}

	rawResp, err := decodeRawResponse[model.Servers](resp)
	if err != nil {
		return nil, rawResp, err
	}
	return *rawResp.Data, rawResp, nil
}

// GetMembers implements CreditPoolAPI.
func (c *creditPoolClient) GetMembers(ctx context.Context) (model.CreditPoolMembers, *RawResponse[model.CreditPoolMembers], error) {
	resp, err := c.do(c.rb.buildCreditPoolMembersGetRequest(
		ctx,
		c.poolID,
	))
	if err != nil {
		return nil, nil, err
	}

	rawResp, err := decodeRawResponse[model.CreditPoolMembers](resp)
	if err != nil {
		return nil, rawResp, err
	}
	return *rawResp.Data, rawResp, nil
}
