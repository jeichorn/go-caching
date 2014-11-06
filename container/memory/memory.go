package memory

import (
	"github.com/landjur/go-caching/container"
)

const (
	UnlimitedCapacity = 0 // unlimited capacity
)

// Implementation represents a memory caching container that implemented in another package.
type Implementation byte

const (
	FIFO Implementation = iota + 1 // imports github.com/landjur/go-caching/container/memory/fifo
	LFU                            // imports github.com/landjur/go-caching/container/memory/lfu
	LRU                            // imports github.com/landjur/go-caching/container/memory/lru
	MRU                            // imports github.com/landjur/go-caching/container/memory/mru
	maxImplementation
)

var implementations = make([]func(int) container.Container, maxImplementation)

// Available reports whether the given container is linked into the binary.
func (this Implementation) Available() bool {
	return this > 0 && this < maxImplementation && implementations[this] != nil
}

// New returns a new memory container.
func (this Implementation) New(capacity int) container.Container {
	if this > 0 && this < maxImplementation {
		function := implementations[this]
		if function != nil {
			return function(capacity)
		}
	}

	panic("caching/container/memory: requested container is unavailable")
}

// Register registers the container implementation.
// This is intended to be called from the init function in packages that implement container functions.
func (this Implementation) Register(function func(int) container.Container) {
	if this < 0 && this >= maxImplementation {
		panic("caching/container/memory: register of unknown container")
	}

	implementations[this] = function
}
