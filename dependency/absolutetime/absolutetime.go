// Copyright 2014 Landjur. All rights reserved.

package absolutetime

import (
	"bitbucket.org/landjur-golang/caching"
	"time"
)

// New returns a new instance of AbsoluteTime.
func New(expirationTime time.Time) *AbsoluteTime {
	return &AbsoluteTime{expirationTime}
}

// AbsoluteTime represents a caching dependency policy by absolute time.
type AbsoluteTime struct {
	ExpirationTime time.Time
}

func (this AbsoluteTime) HasExpired(item *caching.Item) bool {
	return time.Now().After(this.ExpirationTime)
}
