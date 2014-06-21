package inmemory

import (
	"github.com/landjur/go-caching"
	"sync"
	"time"
)

// New returns a new instance of InMemory
func New(capacity int) *InMemory {
	return &InMemory{
		capacity: capacity,
		items:    make(map[string]*caching.Item),
	}
}

// InMemory represents the caching container stores items in-memory.
// InMemory implements caching.Container.
type InMemory struct {
	locker   sync.RWMutex
	capacity int
	items    map[string]*caching.Item
}

// Items returns the items cached.
func (this InMemory) Items() caching.Items {
	this.locker.RLock()
	defer this.locker.RUnlock()

	items := make(caching.Items, len(this.items))
	index := 0
	for _, item := range this.items {
		items[index] = item
		index++
	}

	return items
}

func (this InMemory) Capacity() int {
	return this.capacity
}

func (this InMemory) Count() int {
	return len(this.items)
}

func (this InMemory) Contains(key string) bool {
	_, ok := this.items[key]
	return ok
}

func (this InMemory) Get(key string) (interface{}, error) {
	this.locker.RLock()
	defer this.locker.RUnlock()

	if item, ok := this.items[key]; ok {
		item.LastAccessedTime = time.Now()
		return item.Value, nil
	}

	return nil, nil
}

func (this *InMemory) Set(key string, value interface{}, priority caching.Priority, dependencies ...caching.Dependency) error {
	this.locker.Lock()
	defer this.locker.Unlock()

	//delete(this.items, key)
	item := caching.NewItem(key, value, priority, dependencies...)
	this.items[key] = item
	return nil
}

func (this *InMemory) Remove(key string) error {
	this.locker.Lock()
	defer this.locker.Unlock()

	delete(this.items, key)
	return nil
}

func (this *InMemory) Clear() error {
	this.locker.Lock()
	defer this.locker.Unlock()

	this.items = make(map[string]*caching.Item)
	return nil
}
