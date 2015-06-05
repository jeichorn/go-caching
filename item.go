package caching

import (
	"encoding/gob"
	"time"
)

func init() {
	gob.Register(new(Item))
	gob.Register(time.Time{})
}

// NewItem returns a new instance of Item.
func NewItem(key string, value interface{}, expiration *Expiration, dependencies ...Dependency) *Item {
	return &Item{
		Key:          key,
		Value:        value,
		Expiration:   expiration,
		Dependencies: dependencies,
	}
}

// Item represents an item.
type Item struct {
	Key              string
	Value            interface{}
	CreatedTime      time.Time
	LastAccessedTime time.Time
	Expiration       *Expiration
	Dependencies     []Dependency
}

// HasExpired reports whether the item has expired.
func (this Item) HasExpired() bool {
	if this.Expiration != nil {
		if this.Expiration.AbsoluteTime.Before(time.Now()) || (this.LastAccessedTime.Unix() > 0 && this.LastAccessedTime.Add(this.Expiration.SlidingPeriod).Before(time.Now())) {
			return true
		}
	}

	for _, dependency := range this.Dependencies {
		if dependency != nil {
			if dependency.HasChanged() {
				return true
			}
		}
	}

	return false
}
