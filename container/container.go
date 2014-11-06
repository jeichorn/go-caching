package container

// Container represents the interface of a caching container.
type Container interface {
	// Get returns the value by given key.
	Get(key string) (interface{}, error)

	// Set inserts/updates the item.
	Set(key string, value interface{}) error

	// Remove removes the item by given key.
	Remove(key string) error

	// Clear removes all items.
	Clear() error
}
