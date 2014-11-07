package fifo

import (
	"github.com/landjur/go-caching/container"
	"github.com/landjur/go-caching/container/memory"
)

// New returns a new in-memory caching container using fifo (first in first out) arithmetic.
func New(capacity int) container.Container {
	return &fifo{
		Capacity: capacity,
		Data:     newData(),
	}
}

// register the container.
func init() {
	memory.FIFO.Register(New)
}

type fifo struct {
	Capacity int
	Data     *data
}

func (this *fifo) Get(key string) (interface{}, error) {
	if this.Data == nil {
		return nil, nil
	}

	return this.Data.Get(key), nil
}

func (this *fifo) Set(key string, value interface{}) error {
	if this.Data == nil {
		this.Data = newData()
	}

	if this.Capacity > 0 && this.Data.Count() == this.Capacity && !this.Data.Contains(key) {
		this.Data.Discard()
	}

	this.Data.Set(key, value)

	return nil
}

func (this *fifo) Remove(key string) error {
	if this.Data == nil {
		return nil
	}

	this.Data.Remove(key)

	return nil
}

func (this *fifo) Clear() error {
	if this.Data == nil {
		return nil
	}

	this.Data.Clear()

	return nil
}
