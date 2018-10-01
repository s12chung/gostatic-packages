package atom

import (
	"strings"
	"time"
)

// EntryLimit is the recommended entry limit for the Atom feed
const EntryLimit = 100

// HTMLEntry is an entry for an HTML page
type HTMLEntry struct {
	ID          string
	Title       string
	Updated     time.Time
	HTMLContent string
	Summary     string
	Published   time.Time
}

// ToEntry converts the HTMLEntry to an atom Entry
func (htmlEntry *HTMLEntry) ToEntry(a *Renderer) *Entry {
	return &Entry{
		ID:      strings.Join([]string{a.Settings.Host, htmlEntry.ID, htmlEntry.Published.Format("2006-01-02")}, ":"),
		Title:   htmlEntry.Title,
		Updated: htmlEntry.Updated,

		Author:  a.Author(),
		Content: &EntryContent{Content: htmlEntry.HTMLContent, Type: "html"},
		Summary: htmlEntry.Summary,
		Link:    a.AlternateLink(htmlEntry.ID),

		Published: htmlEntry.Published,
	}
}

// HTMLRenderer is a "main struct", which generates a Feed. It builds on top of Renderer.
type HTMLRenderer struct {
	Settings *Settings
}

// NewHTMLRenderer returns a new instance of HTMLRenderer
func NewHTMLRenderer(settings *Settings) *HTMLRenderer {
	return &HTMLRenderer{settings}
}

// Render  the Atom xml data a generated feed
func (renderer *HTMLRenderer) Render(feedName, selfURL, logoURL string, htmlEntries []*HTMLEntry) ([]byte, error) {
	atomRenderer := NewRenderer(renderer.Settings)
	feed := HTMLEntriesToFeed(atomRenderer, feedName, selfURL, logoURL, htmlEntries)
	return feed.Marhshall()
}

// HTMLEntriesToFeed converts []HTMLEntry to a Feed, with lots of defaults set
func HTMLEntriesToFeed(atomRenderer *Renderer, feedName, selfURL, logoURL string, htmlEntries []*HTMLEntry) *Feed {
	entries := make([]*Entry, len(htmlEntries))
	for i, htmlEntry := range htmlEntries {
		entries[i] = htmlEntry.ToEntry(atomRenderer)
	}

	lastUpdated := time.Now()
	if len(htmlEntries) >= 1 {
		lastUpdated = htmlEntries[0].Updated
	}
	feed := atomRenderer.NewFeed(feedName, lastUpdated, selfURL, logoURL)
	feed.Entries = entries
	return feed
}
