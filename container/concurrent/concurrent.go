package concurrent

import (
	"github.com/landjur/go-caching/container"
	"sync"
)

// New returns a safe container for concurrent access.
func New(real container.Container) container.Container {
	return &concurrent{
		real: real,
	}
}

type concurrent struct {
	locker sync.RWMutex
	real   container.Container
}

func (this *concurrent) Get(key string) (interface{}, error) {
	this.locker.RLock()
	defer this.locker.RUnlock()

	return this.real.Get(key)
}

func (this *concurrent) Set(key string, value interface{}) error {
	this.locker.Lock()
	defer this.locker.Unlock()

	return this.real.Set(key, value)
}

func (this *concurrent) Remove(key string) error {
	this.locker.Lock()
	defer this.locker.Unlock()

	return this.real.Remove(key)
}

func (this *concurrent) Clear() error {
	this.locker.Lock()
	defer this.locker.Unlock()

	return this.real.Clear()
}
