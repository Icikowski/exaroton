package exaroton

import (
	"context"

	"pkg.icikowski.pl/exaroton/model"
)

type playerListClient struct {
	*baseClient

	serverID string
	listName string
}

func newPlayerListClient(base *baseClient, serverID, listName string) PlayerListAPI {
	return &playerListClient{
		baseClient: base,
		serverID:   serverID,
		listName:   listName,
	}
}

// GetList implements PlayerListAPI.
func (c *playerListClient) GetList(ctx context.Context) ([]string, *RawResponse[[]string], error) {
	resp, err := c.do(c.rb.buildPlayerListGetRequest(
		ctx,
		c.serverID,
		c.listName,
	))
	if err != nil {
		return nil, nil, err
	}

	rawResp, err := decodeRawResponse[[]string](resp)
	if err != nil {
		return nil, rawResp, err
	}
	return *rawResp.Data, rawResp, nil
}

// AddPlayers implements PlayerListAPI.
func (c *playerListClient) AddPlayers(ctx context.Context, usernames ...string) ([]string, *RawResponse[[]string], error) {
	resp, err := c.do(c.rb.buildPlayerListPutRequest(
		ctx,
		c.serverID,
		c.listName,
		model.PlayerListEntries{
			Entries: usernames,
		},
	))
	if err != nil {
		return nil, nil, err
	}

	rawResp, err := decodeRawResponse[[]string](resp)
	if err != nil {
		return nil, rawResp, err
	}
	return *rawResp.Data, rawResp, nil
}

// RemovePlayers implements PlayerListAPI.
func (c *playerListClient) RemovePlayers(ctx context.Context, usernames ...string) ([]string, *RawResponse[[]string], error) {
	resp, err := c.do(c.rb.buildPlayerListDeleteRequest(
		ctx,
		c.serverID,
		c.listName,
		model.PlayerListEntries{
			Entries: usernames,
		},
	))
	if err != nil {
		return nil, nil, err
	}

	rawResp, err := decodeRawResponse[[]string](resp)
	if err != nil {
		return nil, rawResp, err
	}
	return *rawResp.Data, rawResp, nil
}
