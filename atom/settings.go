package atom

import (
	"fmt"
	"strings"
)

// Settings presents the settings of the Atom feed
type Settings struct {
	AuthorName string `json:"author_name,omitempty"`
	AuthorURI  string `json:"author_uri,omitempty"`

	Host string `json:"host,omitempty"`
	SSL  bool   `json:"ssl,omitempty"`
}

// DefaultSettings returns the default settings
func DefaultSettings() *Settings {
	return &Settings{
		"Your Name",
		"",
		"yourwebsite.com",
		true,
	}
}

// AuthorURIDefaulted returns a defaulted AuthorURI
func (settings *Settings) AuthorURIDefaulted() string {
	if settings.AuthorURI == "" {
		return settings.URL()
	}
	return settings.AuthorURI
}

// URL returns the home page full URL of your site, including host, protocol, etc.
func (settings *Settings) URL() string {
	ssl := ""
	if settings.SSL {
		ssl = "s"
	}
	return fmt.Sprintf("http%v://%v", ssl, settings.Host)
}

// FullURLFor gives the full URL for the given URL
func (settings *Settings) FullURLFor(url string) string {
	url = strings.Trim(url, "/")
	if url == "" {
		return settings.URL()
	}
	return strings.Join([]string{settings.URL(), url}, "/")
}
