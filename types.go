package exaroton

import (
	"context"
	"io"

	"pkg.icikowski.pl/exaroton/model"
)

// RawResponse represents a raw response from the exaroton's API.
type RawResponse[T any] struct {
	// Success indicates whether the request was successful.
	Success bool `json:"success"`
	// Error contains the error message if the request was not successful.
	Error *string `json:"error"`
	// Data contains the response data.
	Data *T `json:"data"`
}

// PlayerListAPI represents the API for managing player lists.
type PlayerListAPI interface {
	// GetList returns the content of the player list.
	GetList(ctx context.Context) ([]string, *RawResponse[[]string], error)

	// AddPlayers adds players to the player list.
	AddPlayers(ctx context.Context, usernames ...string) ([]string, *RawResponse[[]string], error)
	// RemovePlayers removes players from the player list.
	RemovePlayers(ctx context.Context, usernames ...string) ([]string, *RawResponse[[]string], error)
}

// FileAPI represents the API for managing files.
type FileAPI interface {
	// GetFile returns the file information.
	GetFile(ctx context.Context) (*model.File, *RawResponse[model.File], error)

	// GetContent returns the content of the file.
	GetContent(ctx context.Context, dst io.Writer) (*RawResponse[any], error)
	// SetContent sets the content of the file.
	SetContent(ctx context.Context, src io.Reader) (*RawResponse[any], error)
	// CreateDirectory creates a directory.
	CreateDirectory(ctx context.Context) (*RawResponse[any], error)
	// Delete deletes the file.
	Delete(ctx context.Context) (*RawResponse[any], error)
}

// ConfigAPI represents the API for managing configuration files.
type ConfigAPI interface {
	// GetConfig returns the configuration file information.
	GetConfig(ctx context.Context) (model.ConfigOptions, *RawResponse[model.ConfigOptions], error)
	// SetConfig sets the configuration file data.
	SetConfig(ctx context.Context, opts model.ConfigValues) (model.ConfigOptions, *RawResponse[model.ConfigOptions], error)
}

// ServerAPI represents the API for managing servers.
type ServerAPI interface {
	// GetServer returns the server information.
	GetServer(ctx context.Context) (*model.Server, *RawResponse[model.Server], error)

	// GetLogs returns the server logs.
	GetLogs(ctx context.Context) (*model.ServerLogs, *RawResponse[model.ServerLogs], error)
	// ShareLogs shares the server logs.
	ShareLogs(ctx context.Context) (*model.ServerSharedLogs, *RawResponse[model.ServerSharedLogs], error)

	// GetRAM returns the server's RAM.
	GetRAM(ctx context.Context) (int, *RawResponse[model.ServerRAM], error)
	// SetRAM sets the server's RAM.
	SetRAM(ctx context.Context, gigabytes int) (int, *RawResponse[model.ServerRAM], error)

	// GetMOTD returns the server's MOTD.
	GetMOTD(ctx context.Context) (string, *RawResponse[model.ServerMOTD], error)
	// SetMOTD sets the server's MOTD.
	SetMOTD(ctx context.Context, motd string) (string, *RawResponse[model.ServerMOTD], error)

	// Start starts the server.
	Start(ctx context.Context) (*RawResponse[any], error)
	// StartWithOptions starts the server with the given options.
	StartWithOptions(ctx context.Context, options model.ServerStartParams) (*RawResponse[any], error)
	// Stop stops the server.
	Stop(ctx context.Context) (*RawResponse[any], error)
	// Restart restarts the server.
	Restart(ctx context.Context) (*RawResponse[any], error)
	// Command sends a command to the server.
	Command(ctx context.Context, command string) (*RawResponse[any], error)

	// GetPlayerLists returns the list of player lists.
	GetPlayerLists(ctx context.Context) ([]string, *RawResponse[[]string], error)
	// PlayerList returns the API client for managing single player list.
	PlayerList(name string) PlayerListAPI

	// File returns the API client for managing single file.
	File(path string) FileAPI
	// Config returns the API client for managing single configuration file.
	Config(path string) ConfigAPI
}

// CreditPoolAPI represents the API for managing credit pools.
type CreditPoolAPI interface {
	// GetCreditPool returns the credit pool information.
	GetCreditPool(ctx context.Context) (*model.CreditPool, *RawResponse[model.CreditPool], error)

	// GetServers returns the servers in the credit pool.
	GetServers(ctx context.Context) (model.Servers, *RawResponse[model.Servers], error)
	// GetMembers returns the members of the credit pool.
	GetMembers(ctx context.Context) (model.CreditPoolMembers, *RawResponse[model.CreditPoolMembers], error)
}

// API represents the exaroton's API client.
type API interface {
	// GetAccount returns the account information.
	GetAccount(ctx context.Context) (*model.AccountInfo, *RawResponse[model.AccountInfo], error)

	// GetServers returns the list of available servers.
	GetServers(ctx context.Context) (model.Servers, *RawResponse[model.Servers], error)
	// Server returns the API client for managing single server.
	Server(id string) ServerAPI

	// GetCreditPools returns the list of available credit pools.
	GetCreditPools(ctx context.Context) (model.CreditPools, *RawResponse[model.CreditPools], error)
	// CreditPool returns the API client for managing single credit pool.
	CreditPool(id string) CreditPoolAPI
}
