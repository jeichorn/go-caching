package caching

import (
	"time"
)

// NewExpiration returns a new instance of Expiration.
func NewExpiration(absoluteTime time.Time, slidingPeriod time.Duration) *Expiration {
	return &Expiration{
		AbsoluteTime:  absoluteTime,
		SlidingPeriod: slidingPeriod,
	}
}

// Expiration represents an expiration about time.
type Expiration struct {
	AbsoluteTime  time.Time
	SlidingPeriod time.Duration
}
