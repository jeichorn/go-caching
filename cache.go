package caching

import (
	"fmt"
	"github.com/landjur/go-tasking"
	"reflect"
	"strings"
	"time"
)

// cacheScavengingTask represents a task for scavenging expired tasks.
type cacheScavengingTask struct {
	Container Container
}

// Execute implements tasking.Task interface.
func (this cacheScavengingTask) Execute() {
	// 1. sweep expired items
	items := this.Container.Items()

	var itemsNotExpired Items
	for _, item := range items {
		if item.HasExpired() {
			this.Container.Remove(item.Key)
		} else {
			itemsNotExpired = append(itemsNotExpired, item)
		}
	}

	// 2. sweep scavengable items
	count := len(itemsNotExpired)
	if count > this.Container.Capacity() {
		// Sort the items
		itemsNotExpired.Sort()
		numberToSweep := count - this.Container.Capacity()
		for i := 0; i < numberToSweep; i++ {
			this.Container.Remove(itemsNotExpired[i].Key)
		}
	}
}

// NewCache returns a new instance of Cache.
func NewCache(container Container, scavengingFrequency time.Duration) *Cache {
	cache := &Cache{
		container:           container,
		scavengingScheduler: tasking.NewScheduler(scavengingFrequency, &cacheScavengingTask{container}),
	}

	cache.scavengingScheduler.Start()

	return cache
}

// Cache represents a cache manager object.
type Cache struct {
	container           Container
	scavengingScheduler *tasking.Scheduler
}

// Capacity returns the capacity of cache container.
func (this Cache) Capacity() int {
	return this.container.Capacity()
}

// Count returns the count of items in container.
func (this Cache) Count() int {
	return this.container.Count()
}

// Contains returns true if given key exists.
func (this Cache) Contains(key string) bool {
	return this.container.Contains(key)
}

// Set stores the item.
func (this Cache) Set(key string, value interface{}, priority Priority, dependencies ...Dependency) error {
	if value == nil {
		return fmt.Errorf("caching: no need to cache a nil object(key:%s)", key)
	}

	return this.container.Set(key, value, priority, dependencies...)
}

// Get returns the value cached and true if item key exists, or return nil and false.
func (this Cache) Get(key string, value interface{}) error {
	itemValue, err := this.container.Get(key)
	if err != nil {
		return err
	}

	if itemValue == nil {
		return nil
	}

	target := reflect.ValueOf(value)
	if target.Kind() != reflect.Ptr || target.IsNil() {
		return fmt.Errorf("caching: parsing value of item(key: %s) failed, invalid target value type(%s)", key, reflect.TypeOf(value).String())
	}

	origin := reflect.ValueOf(itemValue)
	target = target.Elem()
	target.Set(origin)

	return nil
}

// Remove removes the item by given key.
func (this Cache) Remove(key string) error {
	return this.container.Remove(key)
}

// RemoveByPrefix removes items key has prefix given.
func (this *Cache) RemoveByPrefix(prefix string) error {
	items := this.container.Items()
	for _, item := range items {
		if strings.HasPrefix(item.Key, prefix) {
			if err := this.Remove(item.Key); err != nil {
				return err
			}
		}
	}

	return nil
}

// Clear removes all items.
func (this Cache) Clear() error {
	return this.container.Clear()
}
