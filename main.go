package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/ikawaha/feedsnippet/feeder"
)

const (
	startTag = `<!--[START github.com/ikawaha/feedsnippet]-->`
	endTag   = `<!--[END github.com/ikawaha/feedsnippet]-->`
)

var snippetField = regexp.MustCompile(`(?s)` + regexp.QuoteMeta(startTag) + `.*` + regexp.QuoteMeta(endTag))

type option struct {
	config  string
	target string
	verbose bool
	debug   bool
}

func main() {
	var opt option
	flag.StringVar(&opt.config, "config", "", "config file")
	flag.StringVar(&opt.target, "file", "", "target file (optional)")
	flag.BoolVar(&opt.verbose, "verbose", false, "print snippets to stdout (optional)")
	flag.BoolVar(&opt.debug, "debug", false, "print raw fees for debug (optional)")
	flag.Parse()
	if err := run(opt); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(opt option) error {
	config, err := feeder.NewConfigFile(opt.config)
	if err != nil {
		return err
	}
	b := bytes.NewBufferString(startTag)
	b.WriteString(fmt.Sprintf("<!--[%s]-->\n", time.Now().Format(time.RFC3339)))
	for _, c := range config {
		f, err := feeder.NewFeeder(feeder.DebugOpt(opt.debug), feeder.FilterOpt(c.Filters()...))
		if err != nil {
			return err
		}
		feeds, err := f.Feeds(c.URLs...)
		if err != nil {
			return fmt.Errorf("feeds error: %w", err)
		}
		tmpl := feeder.DefaultFeedT
		if c.Template != "" {
			tmpl = c.Template
		}
		s, err := feeder.ExecuteTmpl(feeds, tmpl)
		if err != nil {
			return fmt.Errorf("%v\n%s", err, tmpl)
		}
		b.WriteString(s)
	}
	b.WriteString(endTag)
	if err := writeSnippets(opt.target, b.Bytes(), opt.verbose); err != nil {
		return err
	}
	return nil
}

func writeSnippets(dst string, snippet []byte, verbose bool) error {
	if dst == "" || verbose {
		fmt.Println(string(snippet))
		if dst == "" {
			return nil
		}
	}
	body, err := os.ReadFile(dst)
	if err != nil {
		return err
	}
	body = snippetField.ReplaceAll(body, snippet)
	if err := os.WriteFile(dst, body, 0666); err != nil {
		return err
	}
	return nil
}
