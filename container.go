// Copyright 2014 Landjur. All rights reserved.

package caching

// Container represents the interface of container for storing cache items.
type Container interface {
	// Items returns cache items in container.
	Items() Items

	// Capacity returns the capacity of container.
	Capacity() int

	// Count returns the number of items in container current.
	Count() int

	// Contains returns true if item exists by given key.
	Contains(key string) bool

	// Set stores the item to container
	Set(key string, value interface{}, priority Priority, dependencies ...Dependency) error

	// Get parses the item value by given key.
	Get(key string) (interface{}, error)

	// Remove removes the item by given key.
	Remove(key string) error

	// Clear removes all items from container.
	Clear() error
}
