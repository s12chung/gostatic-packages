package markdown

import (
	"path"
	"strings"
	"testing"

	logTest "github.com/sirupsen/logrus/hooks/test"

	"github.com/s12chung/gostatic/go/test"
)

func defaultMarkdown() (*Markdown, *logTest.Hook) {
	log, hook := logTest.NewNullLogger()
	settings := DefaultSettings()
	settings.MarkdownsPath = path.Join(test.FixturePath)
	return NewMarkdown(settings, log), hook
}

func TestMarkdown_ProcessMarkdown(t *testing.T) {
	testCases := []struct {
		filename string
		exp      string
		safeLog  bool
	}{
		{"doesnt_exist.md", "", false},
		{"ProcessMarkdown.md", `<p>Some random <a href="http://stevenchung.ca">markdown</a>.</p>`, true},
	}

	for testCaseIndex, tc := range testCases {
		context := test.NewContext().SetFields(test.ContextFields{
			"index":        testCaseIndex,
			"assetsFolder": tc.filename,
		})

		markdown, hook := defaultMarkdown()
		got := strings.TrimSpace(markdown.ProcessMarkdown(tc.filename))

		if got != tc.exp {
			t.Error(context.GotExpString("Result", got, tc.exp))
		}
		if test.SafeLogEntries(hook) != tc.safeLog {
			t.Error(context.GotExpString("test.SafeLogEntries(hook)", test.SafeLogEntries(hook), tc.safeLog))
		}
	}
}
