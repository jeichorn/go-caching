package mru

import (
	"github.com/landjur/go-caching/container"
	"github.com/landjur/go-caching/container/memory"
)

// New returns a new in-memory caching items using mru (most recently used) arithmetic.
// It is not safe for concurrent access.
func New(capacity int) container.Container {
	return &mru{
		capacity: capacity,
		items:    newItems(),
	}
}

// register the items.
func init() {
	memory.MRU.Register(New)
}

type mru struct {
	capacity int
	items    *items
}

func (this *mru) Get(key string) (interface{}, error) {
	if this.items == nil {
		return nil, nil
	}

	return this.items.Get(key), nil
}

func (this *mru) Set(key string, value interface{}) error {
	if this.items == nil {
		this.items = newItems()
	}

	if this.capacity > 0 && this.items.Count() == this.capacity && !this.items.Contains(key) {
		this.items.Discard()
	}

	this.items.Set(key, value)

	return nil
}

func (this *mru) Remove(key string) error {
	if this.items == nil {
		return nil
	}

	this.items.Remove(key)

	return nil
}

func (this *mru) Clear() error {
	if this.items == nil {
		return nil
	}

	this.items.Clear()

	return nil
}
