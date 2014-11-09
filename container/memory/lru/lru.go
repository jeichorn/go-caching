package lru

import (
	"github.com/landjur/go-caching/container"
	"github.com/landjur/go-caching/container/memory"
)

// New returns a new in-memory caching items using LRU (least recently used) arithmetic.
func New(capacity int) container.Container {
	return &lru{
		capacity: capacity,
		items:    newItems(),
	}
}

// register the items.
func init() {
	memory.LRU.Register(New)
}

type lru struct {
	capacity int
	items    *items
}

func (this *lru) Get(key string) (interface{}, error) {
	if this.items == nil {
		return nil, nil
	}

	return this.items.Get(key), nil
}

func (this *lru) Set(key string, value interface{}) error {
	if this.items == nil {
		this.items = newItems()
	}

	if this.capacity > 0 && this.items.Count() == this.capacity && !this.items.Contains(key) {
		this.items.Discard()
	}

	this.items.Set(key, value)

	return nil
}

func (this *lru) Remove(key string) error {
	if this.items == nil {
		return nil
	}

	this.items.Remove(key)

	return nil
}

func (this *lru) Clear() error {
	if this.items == nil {
		return nil
	}

	this.items.Clear()

	return nil
}
