// Copyright 2014 Landjur. All rights reserved.

package caching

import (
	"sort"
	"time"
)

// NewItem returns a new instance of Item.
func NewItem(key string, value interface{}, priority Priority, dependencies ...Dependency) *Item {
	return &Item{
		Key:          key,
		Value:        value,
		Priority:     priority,
		CreatedTime:  time.Now(),
		Dependencies: dependencies,
	}
}

// Item represents a cache item.
type Item struct {
	Key              string
	Value            interface{}
	Priority         Priority
	CreatedTime      time.Time
	LastAccessedTime time.Time
	Dependencies     []Dependency
}

// HasExpired returns true if cache item expired.
func (this Item) HasExpired() bool {
	for _, dependency := range this.Dependencies {
		if dependency != nil {
			if dependency.HasExpired(&this) {
				return true
			}
		}
	}

	return false
}

// Items represents a collection of Item.
type Items []*Item

// Len implements sort.Interface.
func (this Items) Len() int {
	return len(this)
}

// Less implements sort.Interface.
func (this Items) Less(i, j int) bool {
	if this[i].Priority == this[j].Priority {
		return this[i].LastAccessedTime.Before(this[j].LastAccessedTime)
	}

	return this[i].Priority < this[j].Priority
}

// Swap implements sort.Interface.
func (this Items) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

// Sort sorts the items.
func (this Items) Sort() {
	sort.Sort(this)
}
