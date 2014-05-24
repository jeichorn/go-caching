// Copyright 2014 Landjur. All rights reserved.

package caching

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

// NewCache returns a new instance of Cache.
func NewCache(container Container, scavengingFrequency time.Duration) *Cache {
	cache := &Cache{
		container:           container,
		scavengingFrequency: scavengingFrequency,
	}

	cache.scavengingTimer = time.AfterFunc(cache.scavengingFrequency, func() { cache.scavenging() })

	return cache
}

// Cache represents a cache manager object.
type Cache struct {
	container           Container
	scavengingTimer     *time.Timer
	scavengingFrequency time.Duration
}

// scavenging is the callback function for scavenging cache items.
func (this Cache) scavenging() {
	// 1. sweep expired items
	items := this.container.Items()

	var itemsNotExpired Items
	for _, item := range items {
		if item.HasExpired() {
			this.Remove(item.Key)
		} else {
			itemsNotExpired = append(itemsNotExpired, item)
		}
	}

	// 2. sweep scavengable items
	count := len(itemsNotExpired)
	if count > this.container.Capacity() {
		// Sort the items
		itemsNotExpired.Sort()
		numberToSweep := count - this.container.Capacity()
		for i := 0; i < numberToSweep; i++ {
			this.Remove(itemsNotExpired[i].Key)
		}
	}

	// reset scavenging timer
	if ok := this.scavengingTimer.Reset(this.scavengingFrequency); !ok {
		this.scavengingTimer = time.AfterFunc(this.scavengingFrequency, func() { this.scavenging() })
	}
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
