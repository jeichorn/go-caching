package caching

import (
	"encoding/gob"
	"github.com/landjur/go-caching/dependency"
)

func init() {
	gob.Register(&item{})
}

func newItem(value interface{}, dependencies ...dependency.Dependency) *item {
	return &item{
		Value:        value,
		Dependencies: dependencies,
	}
}

type item struct {
	Value        interface{}
	Dependencies []dependency.Dependency
}

func (this item) HasExpired() bool {
	for _, dependency := range this.Dependencies {
		if dependency != nil {
			if dependency.HasChanged() {
				return true
			}
		}
	}

	return false
}
