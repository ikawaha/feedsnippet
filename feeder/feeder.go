package feeder

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"text/template"

	"github.com/mmcdole/gofeed"
)

var (
	// DefaultFeedT is a default feed template.
	DefaultFeedT = `{{range . -}}
* [{{.Title}}]({{.Link}})
{{end}}`
	defaultTmpl *template.Template
)

func init() {
	var err error
	defaultTmpl, err = template.New("default").Parse(DefaultFeedT)
	if err != nil {
		panic(err)
	}
}

// Feeder reads feeds and create feed snippets from it.
type Feeder struct {
	debug    bool
	debugOut io.Writer
	filters  []Filter
	template string
}

// Option represents a feeder option.
type Option func(*Feeder) error

// FilterOpt is a option setting filters to a feeder.
func FilterOpt(filter ...Filter) Option {
	return func(f *Feeder) error {
		f.filters = append(f.filters, filter...)
		return nil
	}
}

// DebugOpt is a debug option of a feeder.
func DebugOpt(debug bool) Option {
	return func(f *Feeder) error {
		f.debug = debug
		return nil
	}
}

// NewFeeder creates a feeder.
func NewFeeder(opts ...Option) (*Feeder, error) {
	ret := &Feeder{
		debugOut: os.Stderr,
	}
	for _, opt := range opts {
		if err := opt(ret); err != nil {
			return nil, err
		}
	}
	return ret, nil
}

// Feed represents a feed.
type Feed struct {
	Header *gofeed.Feed
	*gofeed.Item
}

// ParseURL fetches the contents of a given url and
// attempts to parse the response into the universal feed type.
func ParseURL(url string) ([]Feed, error) {
	feeds, err := gofeed.NewParser().ParseURL(url)
	if err != nil {
		return nil, fmt.Errorf("parse URL error: %w", err)
	}
	ret := make([]Feed, 0, len(feeds.Items))
	for i := range feeds.Items {
		ret = append(ret, Feed{
			Header: feeds,
			Item:   feeds.Items[i],
		})
	}
	return ret, nil
}

// Feeds fetches the feeds of given urls and applies filters to them.
func (f Feeder) Feeds(urls ...string) ([]Feed, error) {
	var feeds []Feed
	for _, url := range urls {
		fs, err := ParseURL(url)
		if err != nil {
			return nil, err
		}
		if f.debug {
			fmt.Fprintln(f.debugOut, fs)
		}
		feeds = append(feeds, fs...)
	}
	for _, filter := range f.filters {
		var err error
		feeds, err = filter(feeds)
		if err != nil {
			return nil, err
		}
	}
	return feeds, nil
}

// Feeds fetches the feeds of given urls and applies filters to them.
func Feeds(urls []string, filters ...Filter) ([]Feed, error) {
	f, err := NewFeeder(FilterOpt(filters...))
	if err != nil {
		return nil, err
	}
	return f.Feeds(urls...)
}

// ExecuteDefaultTmpl applies a default template to the feeds and returns the result.
func ExecuteDefaultTmpl(feeds []Feed) (string, error) {
	var b bytes.Buffer
	if err := defaultTmpl.Execute(&b, feeds); err != nil {
		return "", err
	}
	return b.String(), nil
}

// ExecuteDefaultTmpl applies the specified template to feeds and returns the result.
func ExecuteTmpl(feeds []Feed, tmpl string) (string, error) {
	t, err := template.New("feed_tmpl").Parse(tmpl)
	if err != nil {
		return "", fmt.Errorf("invalid template: %v", err)
	}
	var b bytes.Buffer
	if err := t.Execute(&b, feeds); err != nil {
		return "", err
	}
	return b.String(), nil
}
