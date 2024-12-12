package model

// File represents the information about server's file.
type File struct {
	// Path represents the path of the file.
	Path string `json:"path"`
	// Name represents the name of the file.
	Name string `json:"name"`
	// IsTextFile represents whether the file is a text file.
	IsTextFile bool `json:"isTextFile"`
	// IsConfigFile represents whether the file is a config file.
	IsConfigFile bool `json:"isConfigFile"`
	// IsDirectory represents whether the file is a directory.
	IsDirectory bool `json:"isDirectory"`
	// IsLog represents whether the file is a log file.
	IsLog bool `json:"isLog"`
	// IsReadable represents whether the file is readable.
	IsReadable bool `json:"isReadable"`
	// IsWritable represents whether the file is writable.
	IsWritable bool `json:"isWritable"`
	// Size represents the size of the file in bytes.
	Size int64 `json:"size"`
	// Children represents the children of the file.
	Children []File `json:"children"`
}

// ConfigFile represents the information about server's config file options.
type ConfigFile struct {
	// Key represents the key of the option.
	Key string `json:"key"`
	// Label represents the label of the option.
	Label string `json:"label"`
	// Type represents the type of the option.
	Type string `json:"type"`
	// Value represents the value of the option.
	Value any `json:"value"`
	// Options represents the list of available option values.
	Options []string `json:"options"`
}
