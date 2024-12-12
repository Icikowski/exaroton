package exaroton

import (
	"context"

	"pkg.icikowski.pl/exaroton/model"
)

type serverClient struct {
	*baseClient

	serverID string
}

func newServerClient(base *baseClient, serverID string) ServerAPI {
	return &serverClient{
		baseClient: base,
		serverID:   serverID,
	}
}

// GetServer implements ServerAPI.
func (c *serverClient) GetServer(ctx context.Context) (*model.Server, *RawResponse[model.Server], error) {
	resp, err := c.do(c.rb.buildServerGetRequest(
		ctx,
		c.serverID,
	))
	if err != nil {
		return nil, nil, err
	}

	rawResp, err := decodeRawResponse[model.Server](resp)
	if err != nil {
		return nil, rawResp, err
	}
	return rawResp.Data, rawResp, nil
}

// GetLogs implements ServerAPI.
func (c *serverClient) GetLogs(ctx context.Context) (*model.ServerLogs, *RawResponse[model.ServerLogs], error) {
	resp, err := c.do(c.rb.buildLogsGetRequest(
		ctx,
		c.serverID,
	))
	if err != nil {
		return nil, nil, err
	}

	rawResp, err := decodeRawResponse[model.ServerLogs](resp)
	if err != nil {
		return nil, rawResp, err
	}
	return rawResp.Data, rawResp, nil
}

// ShareLogs implements ServerAPI.
func (c *serverClient) ShareLogs(ctx context.Context) (*model.ServerSharedLogs, *RawResponse[model.ServerSharedLogs], error) {
	resp, err := c.do(c.rb.buildLogsShareGetRequest(
		ctx,
		c.serverID,
	))
	if err != nil {
		return nil, nil, err
	}

	rawResp, err := decodeRawResponse[model.ServerSharedLogs](resp)
	if err != nil {
		return nil, rawResp, err
	}
	return rawResp.Data, rawResp, nil
}

// GetRAM implements ServerAPI.
func (c *serverClient) GetRAM(ctx context.Context) (int, *RawResponse[model.ServerRAM], error) {
	resp, err := c.do(c.rb.buildRAMGetRequest(
		ctx,
		c.serverID,
	))
	if err != nil {
		return 0, nil, err
	}

	rawResp, err := decodeRawResponse[model.ServerRAM](resp)
	if err != nil {
		return 0, rawResp, err
	}
	return rawResp.Data.RAM, rawResp, nil
}

// SetRAM implements ServerAPI.
func (c *serverClient) SetRAM(ctx context.Context, gigabytes int) (int, *RawResponse[model.ServerRAM], error) {
	resp, err := c.do(c.rb.buildRAMPostRequest(
		ctx,
		c.serverID,
		model.ServerRAM{
			RAM: gigabytes,
		},
	))
	if err != nil {
		return 0, nil, err
	}

	rawResp, err := decodeRawResponse[model.ServerRAM](resp)
	if err != nil {
		return 0, rawResp, err
	}
	return rawResp.Data.RAM, rawResp, nil
}

// GetMOTD implements ServerAPI.
func (c *serverClient) GetMOTD(ctx context.Context) (string, *RawResponse[model.ServerMOTD], error) {
	resp, err := c.do(c.rb.buildMOTDGetRequest(
		ctx,
		c.serverID,
	))
	if err != nil {
		return "", nil, err
	}

	rawResp, err := decodeRawResponse[model.ServerMOTD](resp)
	if err != nil {
		return "", rawResp, err
	}
	return rawResp.Data.MOTD, rawResp, nil
}

// SetMOTD implements ServerAPI.
func (c *serverClient) SetMOTD(ctx context.Context, motd string) (string, *RawResponse[model.ServerMOTD], error) {
	resp, err := c.do(c.rb.buildMOTDPostRequest(
		ctx,
		c.serverID,
		model.ServerMOTD{
			MOTD: motd,
		},
	))
	if err != nil {
		return "", nil, err
	}

	rawResp, err := decodeRawResponse[model.ServerMOTD](resp)
	if err != nil {
		return "", rawResp, err
	}
	return rawResp.Data.MOTD, rawResp, nil
}

// Start implements ServerAPI.
func (c *serverClient) Start(ctx context.Context) (*RawResponse[any], error) {
	resp, err := c.do(c.rb.buildActionStartGetRequest(
		ctx,
		c.serverID,
	))
	if err != nil {
		return nil, err
	}

	rawResp, err := decodeRawResponse[any](resp)
	if err != nil {
		return nil, err
	}
	return rawResp, nil
}

// StartWithOptions implements ServerAPI.
func (c *serverClient) StartWithOptions(ctx context.Context, options model.ServerStartParams) (*RawResponse[any], error) {
	resp, err := c.do(c.rb.buildActionStartPostRequest(
		ctx,
		c.serverID,
		options,
	))
	if err != nil {
		return nil, err
	}

	rawResp, err := decodeRawResponse[any](resp)
	if err != nil {
		return nil, err
	}
	return rawResp, nil
}

// Stop implements ServerAPI.
func (c *serverClient) Stop(ctx context.Context) (*RawResponse[any], error) {
	resp, err := c.do(c.rb.buildActionStopGetRequest(
		ctx,
		c.serverID,
	))
	if err != nil {
		return nil, err
	}

	rawResp, err := decodeRawResponse[any](resp)
	if err != nil {
		return nil, err
	}
	return rawResp, nil
}

// Restart implements ServerAPI.
func (c *serverClient) Restart(ctx context.Context) (*RawResponse[any], error) {
	resp, err := c.do(c.rb.buildActionRestartGetRequest(
		ctx,
		c.serverID,
	))
	if err != nil {
		return nil, err
	}

	rawResp, err := decodeRawResponse[any](resp)
	if err != nil {
		return nil, err
	}
	return rawResp, nil
}

// Command implements ServerAPI.
func (c *serverClient) Command(ctx context.Context, command string) (*RawResponse[any], error) {
	resp, err := c.do(c.rb.buildActionCommandPostRequest(
		ctx,
		c.serverID,
		model.ServerCommand{
			Command: command,
		},
	))
	if err != nil {
		return nil, err
	}

	rawResp, err := decodeRawResponse[any](resp)
	if err != nil {
		return nil, err
	}
	return rawResp, nil
}

// GetPlayerLists implements ServerAPI.
func (c *serverClient) GetPlayerLists(ctx context.Context) ([]string, *RawResponse[[]string], error) {
	resp, err := c.do(c.rb.buildPlayerListsGetRequest(
		ctx,
		c.serverID,
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

// PlayerList implements ServerAPI.
func (c *serverClient) PlayerList(name string) PlayerListAPI {
	return newPlayerListClient(c.baseClient, c.serverID, name)
}

// File implements ServerAPI.
func (c *serverClient) File(path string) FileAPI {
	return newFileClient(c.baseClient, c.serverID, path)
}

// Config implements ServerAPI.
func (c *serverClient) Config(path string) ConfigAPI {
	return newConfigClient(c.baseClient, c.serverID, path)
}
