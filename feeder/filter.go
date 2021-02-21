package feeder

import (
	"sort"
	"time"
)

// Filter filters feeds.
type Filter func(feed []Feed) ([]Feed, error)

// SortByPublished sorts feeds by published, most recent first.
func SortByPublished() Filter {
	return func(feed []Feed) ([]Feed, error) {
		sort.Slice(feed, func(i, j int) bool {
			var x, y time.Time
			if feed[i].PublishedParsed != nil {
				x = *feed[i].PublishedParsed
			}
			if feed[j].PublishedParsed != nil {
				y = *feed[j].PublishedParsed
			}
			return x.After(y) // note: most recent first.
		})
		return feed, nil
	}
}

// Reverse sorts feeds by reverse order.
func Reverse() Filter {
	return func(feed []Feed) ([]Feed, error) {
		for i := 0; i < len(feed)/2; i++ {
			feed[i], feed[len(feed)-i-1] = feed[len(feed)-i-1], feed[i]
		}
		return feed, nil
	}
}

// Limit limits the number of feeds to n.
func Limit(n int) Filter {
	return func(feed []Feed) ([]Feed, error) {
		if size := len(feed); size < n {
			n = size
		}
		feed = feed[:n]
		return feed, nil
	}
}
