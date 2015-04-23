#Go-Caching

The go-caching packaging providers a scalable caching component for golang. 

##Container

The container represents an adapter for caching manager/service.

###Available Containers

- Concurrent Container: wrapping a container for concurrent access.

- Multi-Level Container: wrapping the containers into a single container.

- Memory Containers: local memory containers (not safe for concurrent access).

    - FIFO Container: replacement algorithm using FIFO (first in first out).

    - LFU Container: replacement algorithm using LFU (least frequently used).

    - LRU Container: replacement algorithm using LRU (least recently used).

    - MRU Container: replacement algorithm using MRU (most recently used).

    - ARC Container: replacement algorithm using ARC (adaptive/adjustable replacement cache).

- Memcache Container: the adapter for [memcached](http://memcached.org).

##Dependency

The dependency represents an external expiration policy.

###Available Dependencies

- File Dependency: it's very useful for caching configuration file.

##Examples

###Example: File Configuration Caching

        import (
            "github.com/wayn3h0/go-caching"
            "github.com/wayn3h0/go-caching/container/memory"
            _ "github.com/wayn3h0/go-caching/container/memory/arc"
            "github.com/wayn3h0/go-caching/dependency/file"
        )

        container := memory.ARC.New(1000)                           // local memory container (ARC), capacity: 1000 (NOTE: not safe for concurrent access)
        cache := caching.NewCache(container)
        cache.Set("key", "value")                                   // memory container always returns nil error
        cache.Set("configuration-key", "settings", file.New(path))  // dependency by configuration file

###Example: Multi-Level Caching 

        import (
            "github.com/wayn3h0/go-caching"
            "github.com/wayn3h0/go-caching/container/concurrent"
            "github.com/wayn3h0/go-caching/container/memcache"
            "github.com/wayn3h0/go-caching/container/memory"
            _ "github.com/wayn3h0/go-caching/container/memory/arc"
            "github.com/wayn3h0/go-caching/container/multilevel"
        )

        level1 := concurrent.New(memory.ARC.New(1000))              // local memory container (ARC), capacity: 1000, safe for concurrent access
        level2 := memcache.New("192.168.100.101:11211")
        container := multilevel.New(level1, level2)                 // local memory as level 1, memcached as level 2
        cache := caching.NewCache(container)
        err := cache.Set("key", "value")                            // memcached may return error
        if err != nil {
            // handle error
        }

