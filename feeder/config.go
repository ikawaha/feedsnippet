package feeder

import (
	"fmt"
	"io"
	"os"

	"github.com/goccy/go-yaml"
)

// NewConfigFile creates a feeder config from a file.
func NewConfigFile(path string) ([]Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	b, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	var config []Config
	if err := yaml.Unmarshal(b, &config); err != nil {
		return nil, fmt.Errorf("invalid config: %v", err)
	}
	return config, nil
}

// Config is a config of a feeder.
type Config struct {
	URLs     []string `json:"urls"`
	Template string   `json:"template,omitempty"`

	// Filters
	Limit           int  `json:"limit,omitempty"`
	SortByPublished bool `json:"sort_by_published,omitempty"`
	Reverse         bool `json:"reverse,omitempty"`
}

// Filters returns filters specified in a config.
func (c Config) Filters() []Filter {
	var ret []Filter
	if c.SortByPublished {
		ret = append(ret, SortByPublished())
	}
	if c.Reverse {
		ret = append(ret, Reverse())
	}
	if c.Limit > 0 {
		ret = append(ret, Limit(c.Limit))
	}
	return ret
}
