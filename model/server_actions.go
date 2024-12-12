package model

// ServerStartParams represents the information required for starting server.
type ServerStartParams struct {
	// UseOwnCredits represents whether to use own credits for starting server.
	UseOwnCredits bool `json:"useOwnCredits"`
}

// ServerCommand represents the information required for sending command to server.
type ServerCommand struct {
	// Command represents the command to be sent to server.
	Command string `json:"command"`
}
