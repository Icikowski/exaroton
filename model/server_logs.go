package model

// ServerLogs represents the logs of a server.
type ServerLogs struct {
	// Logs represents the content of the logs.
	Content *string `json:"content"`
}

// ServerSharedLogs represents the shared log details.
type ServerSharedLogs struct {
	// ID represents the ID of the shared log on mclo.gs.
	ID string `json:"id"`
	// URL represents the URL of the shared log on mclo.gs.
	URL string `json:"url"`
	// Raw represents the URL o API of the shared log on mclo.gs.
	Raw string `json:"raw"`
}
