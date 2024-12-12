package model

// PlayerListEntries represents the information about server's player list entries.
type PlayerListEntries struct {
	// Entries represents the list of players' usernames.
	Entries []string `json:"entries"`
}
