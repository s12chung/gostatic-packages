package goodreads

// Settings represents the settings of the Client
type Settings struct {
	CachePath  string `json:"cache_path,omitempty"`
	APIURL     string `json:"api_url,omitempty"`
	APIKey     string `json:"api_key,omitempty"`
	UserID     int    `json:"user_id,omitempty"`
	PerPage    int    `json:"per_page,omitempty"`
	MaxPerPage int    `json:"max_per_page,omitempty"`
	RateLimit  int    `json:"rate_limit,omitempty"`
}

func (settings *Settings) invalid() bool {
	return settings.APIKey == "" && settings.UserID == 0
}

// DefaultSettings represnts the default Settings
func DefaultSettings() *Settings {
	return &Settings{
		"./cache",
		"https://www.goodreads.com",
		"",
		0,
		50,
		200,
		1000,
	}
}
