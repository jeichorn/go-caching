// Copyright 2014 Landjur. All rights reserved.

package caching

// Dependency represents the interface of an expires policy for cache item.
type Dependency interface {
	HasExpired(item *Item) bool
}
