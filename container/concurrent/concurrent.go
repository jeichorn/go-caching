package concurrent

import (
	"github.com/wayn3h0/go-caching"
	"sync"
)

// New returns a new instance of caching.Container: safe container wrapper for container access.
func New(real caching.Container) caching.Container {
	return &container{
		real: real,
	}
}

type container struct {
	locker sync.RWMutex
	real   caching.Container
}

func (this *container) Get(key string) (interface{}, error) {
	this.locker.RLock()
	defer this.locker.RUnlock()

	return this.real.Get(key)
}

func (this *container) Set(key string, value interface{}) error {
	this.locker.Lock()
	defer this.locker.Unlock()

	return this.real.Set(key, value)
}

func (this *container) Remove(key string) error {
	this.locker.Lock()
	defer this.locker.Unlock()

	return this.real.Remove(key)
}

func (this *container) Clear() error {
	this.locker.Lock()
	defer this.locker.Unlock()

	return this.real.Clear()
}
