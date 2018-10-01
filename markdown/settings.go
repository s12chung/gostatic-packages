package markdown

// Settings contains the settings for the Markdown
type Settings struct {
	MarkdownsPath string `json:"path,omitempty"`
}

// DefaultSettings returns the default Settings
func DefaultSettings() *Settings {
	return &Settings{
		"./content/markdowns",
	}
}
