package lfu

import (
	c "github.com/landjur/go-caching/container"
	"github.com/landjur/go-caching/container/memory"
)

// New returns a new in-memory caching container using lfu (least frequently used) arithmetic.
func New(capacity int) c.Container {
	return &lfu{
		container: newContainer(),
		Capacity:  capacity,
	}
}

// register the container.
func init() {
	memory.LFU.Register(New)
}

type lfu struct {
	container *container
	Capacity  int
}

func (this *lfu) Get(key string) (interface{}, error) {
	if this.container == nil {
		return nil, nil
	}

	return this.container.Get(key), nil
}

func (this *lfu) Set(key string, value interface{}) error {
	if this.container == nil {
		this.container = newContainer()
	}

	if this.Capacity > 0 && this.container.Count() == this.Capacity && !this.container.Contains(key) {
		this.container.Discard()
	}

	this.container.Set(key, value)

	return nil
}

func (this *lfu) Remove(key string) error {
	if this.container == nil {
		return nil
	}

	this.container.Remove(key)

	return nil
}

func (this *lfu) Clear() error {
	if this.container == nil {
		return nil
	}

	this.container.Clear()

	return nil
}
