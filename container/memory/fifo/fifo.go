package fifo

import (
	"github.com/landjur/go-caching/container"
	"github.com/landjur/go-caching/container/memory"
)

// New returns a new in-memory caching container using fifo (first in first out) arithmetic.
func New(capacity int) container.Container {
	return &fifo{
		capacity: capacity,
		items:    newItems(),
	}
}

// register the container.
func init() {
	memory.FIFO.Register(New)
}

type fifo struct {
	capacity int
	items    *items
}

func (this *fifo) Get(key string) (interface{}, error) {
	if this.items == nil {
		return nil, nil
	}

	return this.items.Get(key), nil
}

func (this *fifo) Set(key string, value interface{}) error {
	if this.items == nil {
		this.items = newItems()
	}

	if this.capacity > 0 && this.items.Count() == this.capacity && !this.items.Contains(key) {
		this.items.Discard()
	}

	this.items.Set(key, value)

	return nil
}

func (this *fifo) Remove(key string) error {
	if this.items == nil {
		return nil
	}

	this.items.Remove(key)

	return nil
}

func (this *fifo) Clear() error {
	if this.items == nil {
		return nil
	}

	this.items.Clear()

	return nil
}
