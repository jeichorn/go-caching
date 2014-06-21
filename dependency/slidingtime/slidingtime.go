package slidingtime

import (
	"github.com/landjur/go-caching"
	"time"
)

// New returns a new instance of SlidingTime.
func New(period time.Duration) *SlidingTime {
	return &SlidingTime{period}
}

// SlidingTime represents a caching dependency policy by sliding time.
type SlidingTime struct {
	Period time.Duration
}

func (this SlidingTime) HasExpired(item *caching.Item) bool {
	return time.Now().After(item.LastAccessedTime.Add(this.Period))
}
