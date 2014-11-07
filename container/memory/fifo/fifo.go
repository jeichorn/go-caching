package fifo

import (
	c "github.com/landjur/go-caching/container"
	"github.com/landjur/go-caching/container/memory"
)

// New returns a new in-memory caching container using fifo (first in first out) arithmetic.
func New(capacity int) c.Container {
	return &fifo{
		container: newContainer(),
		Capacity:  capacity,
	}
}

// register the container.
func init() {
	memory.FIFO.Register(New)
}

type fifo struct {
	container *container
	Capacity  int
}

func (this *fifo) Get(key string) (interface{}, error) {
	if this.container == nil {
		return nil, nil
	}

	return this.container.Get(key), nil
}

func (this *fifo) Set(key string, value interface{}) error {
	if this.container == nil {
		this.container = newContainer()
	}

	if this.Capacity > 0 && this.container.Count() == this.Capacity && !this.container.Contains(key) {
		this.container.Discard()
	}

	this.container.Set(key, value)

	return nil
}

func (this *fifo) Remove(key string) error {
	if this.container == nil {
		return nil
	}

	this.container.Remove(key)

	return nil
}

func (this *fifo) Clear() error {
	if this.container == nil {
		return nil
	}

	this.container.Clear()

	return nil
}
