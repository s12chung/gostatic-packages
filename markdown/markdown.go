/*
Package markdown is a basic github.com/s12chung/gostatic/go/lib/router/html.Plugin showing markdown HTMl in templates.
*/
package markdown

import (
	"html/template"
	"io/ioutil"
	"path"

	"github.com/russross/blackfriday"
	"github.com/sirupsen/logrus"
)

// Markdown is a "main struct", which has the configuration to find the markdown files
type Markdown struct {
	settings *Settings
	log      logrus.FieldLogger
}

// NewMarkdown returns a new instance of Markdown
func NewMarkdown(settings *Settings, log logrus.FieldLogger) *Markdown {
	return &Markdown{settings, log}
}

// ProcessMarkdown returns the HTML of the markdown of the given filepath relative to
// Settings.MarkdownsPath
func (markdown *Markdown) ProcessMarkdown(filepath string) string {
	filePath := path.Join(markdown.settings.MarkdownsPath, filepath)
	input, err := ioutil.ReadFile(filePath)
	if err != nil {
		markdown.log.Error(err)
		return ""
	}
	return string(blackfriday.Run(input))
}

// TemplateFuncs is the list of functions provided to the HTML templates
func (markdown *Markdown) TemplateFuncs() template.FuncMap {
	return template.FuncMap{
		"markdown": markdown.ProcessMarkdown,
	}
}
