package caching

import (
	"github.com/landjur/go-caching/container"
	"github.com/landjur/go-caching/dependency"
)

// New returns a new cache.
func New(container container.Container) *Cache {
	return &Cache{
		container: container,
	}
}

// Cache represents a cache.
type Cache struct {
	container container.Container
}

// Get returns the item by given key.
func (this *Cache) Get(key string) (interface{}, error) {
	if this.container == nil {
		return nil, nil
	}

	value, err := this.container.Get(key)
	if err != nil {
		return nil, err
	}
	item := value.(*item)
	if item.HasExpired() {
		err = this.container.Remove(key)
		return nil, err
	}

	return item.Value, nil
}

// Set sets(inserts/updates) item.
func (this Cache) Set(key string, value interface{}, dependencies ...dependency.Dependency) error {
	if this.container == nil {
		return nil
	}

	return this.container.Set(key, newItem(value, dependencies...))
}

// Remove removes the item by given key.
func (this Cache) Remove(key string) error {
	if this.container == nil {
		return nil
	}

	return this.container.Remove(key)
}

// Clear removes all items.
func (this Cache) Clear() error {
	return this.container.Clear()
}
