package caching

// NewCache returns a new instance of Cache.
func NewCache(container Container) *Cache {
	return &Cache{
		container: container,
	}
}

// Cache represents a cache.
type Cache struct {
	container Container
}

// Get returns the value of caching item. it returns nil if caching item has expired.
func (this *Cache) Get(key string) (interface{}, error) {
	if this.container == nil {
		return nil, nil
	}

	value, err := this.container.Get(key)
	if err != nil {
		return nil, err
	}
	item := value.(*Item)
	if item.HasExpired() {
		err = this.container.Remove(key)
		return nil, err
	}

	return item.Value, nil
}

// Set sets the caching item.
func (this Cache) Set(key string, value interface{}, expiration *Expiration, dependencies ...Dependency) error {
	if this.container == nil {
		return nil
	}

	return this.container.Set(key, NewItem(key, value, expiration, dependencies...))
}

// Remove removes the caching item.
func (this Cache) Remove(key string) error {
	if this.container == nil {
		return nil
	}

	return this.container.Remove(key)
}

// Clear removes all caching items.
func (this Cache) Clear() error {
	return this.container.Clear()
}
