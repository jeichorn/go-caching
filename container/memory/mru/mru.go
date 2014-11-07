package mru

import (
	c "github.com/landjur/go-caching/container"
	"github.com/landjur/go-caching/container/memory"
)

// New returns a new in-memory caching container using mru (most recently used) arithmetic.
// It is not safe for concurrent access.
func New(capacity int) c.Container {
	return &mru{
		container: newContainer(),
		Capacity:  capacity,
	}
}

// register the container.
func init() {
	memory.MRU.Register(New)
}

type mru struct {
	container *container
	Capacity  int
}

func (this *mru) Get(key string) (interface{}, error) {
	if this.container == nil {
		return nil, nil
	}

	return this.container.Get(key), nil
}

func (this *mru) Set(key string, value interface{}) error {
	if this.container == nil {
		this.container = newContainer()
	}

	if this.Capacity > 0 && this.container.Count() == this.Capacity && !this.container.Contains(key) {
		this.container.Discard()
	}

	this.container.Set(key, value)

	return nil
}

func (this *mru) Remove(key string) error {
	if this.container == nil {
		return nil
	}

	this.container.Remove(key)

	return nil
}

func (this *mru) Clear() error {
	if this.container == nil {
		return nil
	}

	this.container.Clear()

	return nil
}
