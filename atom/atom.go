/*
Package atom represents your Atom xml data into structs and Marshalls them.
*/
package atom

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"
)

// Feed represents the entire Atom feed
type Feed struct {
	XMLName xml.Name `xml:"feed"`

	XMLLang string `xml:"xml:lang,attr"`
	XMLNS   string `xml:"xmlns,attr"`

	ID      string    `xml:"id"`
	Title   string    `xml:"title"`
	Updated time.Time `xml:"updated"`

	Icon   string  `xml:"icon"`
	Author *Author `xml:"author"`

	Links   []*Link  `xml:"link"`
	Entries []*Entry `xml:"entry"`
}

// Marhshall returns the Atom xml data of the feed
func (feed *Feed) Marhshall() ([]byte, error) {
	bytes, err := xml.MarshalIndent(&feed, "", "  ")
	if err != nil {
		return nil, err
	}
	return append([]byte(xml.Header), bytes...), nil
}

// Entry represents an entry in the Atom feed
type Entry struct {
	XMLName xml.Name `xml:"entry"`

	ID      string    `xml:"id"`
	Title   string    `xml:"title"`
	Updated time.Time `xml:"updated"`

	Author  *Author       `xml:"author"`
	Content *EntryContent `xml:"content"`
	Link    *Link         `xml:"link"`
	Summary string        `xml:"summary"`

	Published time.Time `xml:"published"`
}

// Author represents an author in the Atom feed
type Author struct {
	XMLName xml.Name `xml:"author"`

	Name string `xml:"name,omitempty"`
	URI  string `xml:"uri,omitempty"`
}

// EntryContent represents the content of a Entry in the Atom feed
type EntryContent struct {
	XMLName xml.Name `xml:"content"`

	Content string `xml:",cdata"`
	Type    string `xml:"type,attr,omitempty"`
}

// Link represents a link in the Atom feed
type Link struct {
	XMLName xml.Name `xml:"link"`

	Rel  string `xml:"rel,attr,omitempty"`
	Type string `xml:"type,attr,omitempty"`
	Href string `xml:"href,attr"`
}

// Renderer is a "main struct", which generates a Feed
type Renderer struct {
	Settings *Settings
}

// NewRenderer returns a new instance of Renderer
func NewRenderer(settings *Settings) *Renderer {
	return &Renderer{settings}
}

// Author returns the Author of the feed
func (a *Renderer) Author() *Author {
	return &Author{Name: a.Settings.AuthorName, URI: a.Settings.AuthorURIDefaulted()}
}

// AlternateLink returns the HTML page variation of the feed
func (a *Renderer) AlternateLink(url string) *Link {
	return &Link{Rel: "alternate", Type: "text/html", Href: a.Settings.FullURLFor(url)}
}

// NewFeed returns a new Feed, with lots of defaults set
func (a *Renderer) NewFeed(feedName string, lastUpdated time.Time, selfURL, iconURL string) *Feed {
	return &Feed{
		XMLLang: "en-US",
		XMLNS:   "http://www.w3.org/2005/Atom",

		Title:   fmt.Sprintf("%v - %v", strings.Title(feedName), a.Settings.Host),
		Icon:    a.Settings.FullURLFor(iconURL),
		ID:      strings.Join([]string{a.Settings.Host, "2018", feedName}, ":"),
		Updated: lastUpdated,

		Author: a.Author(),

		Links: []*Link{
			{Rel: "self", Type: "application/atom+xml", Href: a.Settings.FullURLFor(selfURL)},
			a.AlternateLink(""),
		},
	}
}
