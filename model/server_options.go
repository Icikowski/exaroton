package model

// ServerRAM represents the information about server's ServerRAM.
type ServerRAM struct {
	// RAM represents the amount of server's RAM in gigabytes.
	RAM int `json:"ram"`
}

// ServerMOTD represents the information about server's ServerMOTD.
type ServerMOTD struct {
	// MOTD represents the MOTD of the server.
	MOTD string `json:"motd"`
}
