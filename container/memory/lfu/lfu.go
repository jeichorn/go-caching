package lfu

import (
	"github.com/wayn3h0/go-caching"
	"github.com/wayn3h0/go-caching/container/memory"
)

// New returns a new instance of caching.Container: in-memory caching container using lfu (least frequently used) arithmetic.
func New(capacity int) caching.Container {
	return &container{
		capacity: capacity,
		items:    newItems(),
	}
}

// register the items.
func init() {
	memory.LFU.Register(New)
}

type container struct {
	capacity int
	items    *items
}

func (this *container) Get(key string) (interface{}, error) {
	if this.items == nil {
		return nil, nil
	}

	return this.items.Get(key), nil
}

func (this *container) Set(key string, value interface{}) error {
	if this.items == nil {
		this.items = newItems()
	}

	if this.capacity > 0 && this.items.Count() == this.capacity && !this.items.Contains(key) {
		this.items.Discard()
	}

	this.items.Set(key, value)

	return nil
}

func (this *container) Remove(key string) error {
	if this.items == nil {
		return nil
	}

	this.items.Remove(key)

	return nil
}

func (this *container) Clear() error {
	if this.items == nil {
		return nil
	}

	this.items.Clear()

	return nil
}
