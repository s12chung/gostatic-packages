# gostatic-packages
[![GoDoc](https://godoc.org/github.com/s12chung/gostatic-packages?status.svg)](https://godoc.org/github.com/s12chung/gostatic-packages)
[![Build Status](https://travis-ci.com/s12chung/gostatic-packages.svg?branch=master)](https://travis-ci.com/s12chung/gostatic-packages)
[![Coverage Status](https://coveralls.io/repos/github/s12chung/gostatic-packages/badge.svg?branch=master)](https://coveralls.io/github/s12chung/gostatic-packages?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/s12chung/gostatic-packages)](https://goreportcard.com/report/github.com/s12chung/gostatic-packages)

Set of packages to use with [gostatic](https://github.com/s12chung/gostatic).

## Basic Packages

Basic Go packages.

- `robots` - generates a robots.txt by representing them with structs

## Plugin Packages

Plugin packages are packages that implement [gostatic/lib/html.Plugin](https://godoc.org/github.com/s12chung/gostatic/go/lib/html#Plugin)
to access functions in Go HTML templates. You may need to configure the package with a `Settings` struct.
After, create an instance of the "main struct" (which implements the `Plugin` interface) and add that to the plugins for the HTML renderer.

- `markdown` - `Markdown` allows the HTML template to output markdown file HTML

For example:
```go
// In https://github.com/s12chung/gostatic/blob/master/blueprint/go/content/settings.go,
// add another setting:

type Settings struct {
	HTML    *html.Settings    `json:"html,omitempty"`
	Webpack *webpack.Settings `json:"webpack,omitempty"`
	// Newly added Setting
	Markdown *markdown.Settings `json:"markdown,omitempty"`
}


// Settings.json can configure the package set from
// https://github.com/s12chung/gostatic/blob/master/blueprint/main.go
// 
// Pass those settings to the "main struct" in
// https://github.com/s12chung/gostatic/blob/master/blueprint/go/content/content.go
// and add it to the HTML renderer:

func NewContent(generatedPath string, settings *Settings, log logrus.FieldLogger) *Content {
	w := webpack.NewWebpack(generatedPath, settings.Webpack, log)
	// The "main struct"
	md := markdown.NewMarkdown(settings.Markdown)
	htmlRenderer := html.NewRenderer(settings.HTML, []html.Plugin{w, md}, log)
	return &Content{settings, log, htmlRenderer, w, atomRenderer}
}

// Now you can call `Markdown`'s `TemplateFuncs` in Go HTML templates via `html.Renderer`
```


## Settings Packages
Like Plugin Packages, Settings Packages are configured with a `Settings` struct. After, you create an instance of a "main struct"
(there can be more than 1 "main struct") that takes a `Settings` struct and use it in the route.

- `atom` - `atom.Renderer` and `atom.HTMLRenderer` generates an Atom feed given their respective entries
- `goodreads` - `goodreads.Client` retrieves books and reviews from the Goodreads API

For example:

```go
// In https://github.com/s12chung/gostatic/blob/master/blueprint/go/content/settings.go,
// add another setting:

type Settings struct {
	HTML    *html.Settings    `json:"html,omitempty"`
	Webpack *webpack.Settings `json:"webpack,omitempty"`
	// Newly added Setting
	Atom *atom.Settings `json:"atom,omitempty"`
}

// Settings.json can configure the package set from
// https://github.com/s12chung/gostatic/blob/master/blueprint/main.go
// 
// Pass those settings to the "main struct" in
// https://github.com/s12chung/gostatic/blob/master/blueprint/go/content/content.go

type Content struct {
	Settings *Settings
	Log      logrus.FieldLogger

	HTMLRenderer *html.Renderer
	Webpack      *webpack.Webpack
	
	// Newly added so the route can access it
	AtomRenderer *atom.HtmlRenderer
}

func NewContent(generatedPath string, settings *Settings, log logrus.FieldLogger) *Content {
	w := webpack.NewWebpack(generatedPath, settings.Webpack, log)
	htmlRenderer := html.NewRenderer(settings.HTML, []html.Plugin{w}, log)
	// The "main struct"
	atomRenderer := atom.NewHtmlRenderer(settings.Atom)
	return &Content{settings, log, htmlRenderer, w, atomRenderer}
}

// Then you can use the package in the route:

func (content *Content) SetRoutes(r router.Router, tracker *app.Tracker) {
	r.GetRootHTML(content.getRoot)
	r.GetHTML("/404.html", content.get404)
	r.GetHTML("/robots.txt", content.getRobots)
	
	// The new route
	r.Get("/posts.atom", content.getPostsAtom)
}

func (content *Content) getPostsAtom(ctx router.Context) error {
	posts := findPosts()
	
	logoUrl := content.Webpack.ManifestUrl("images/logo.png")
	htmlEntries := postsToHtmlEntries(posts)
	
	bytes, err := content.AtomRenderer.Render("someFeedName", ctx.Url(), logoUrl, htmlEntries)
	if err != nil {
		return err
	}
	return ctx.Respond(bytes)
}
```