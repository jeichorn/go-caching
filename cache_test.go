package caching

import (
	"fmt"
	"github.com/landjur/go-caching/container/concurrent"
	"github.com/landjur/go-caching/container/memcache"
	"github.com/landjur/go-caching/container/memory"
	_ "github.com/landjur/go-caching/container/memory/arc"
	"github.com/landjur/go-caching/container/multilevel"
	"github.com/landjur/go-caching/dependency/file"
	"testing"
)

func Test(t *testing.T) {
	level1 := concurrent.New(memory.ARC.New(1000)) // local memory container (ARC), capacity: 1000, safe for concurrent access
	level2 := memcache.New("192.168.100.101:11211")
	container := multilevel.New(level1, level2) // local memory as level 1, memcached as level 2
	cache := New(container)
	err := cache.Set("key", "value") // memcached may return error
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 2000; i++ {
		if i%2 == 0 {
			err := cache.Set(fmt.Sprintf("%d", i), i)
			if err != nil {
				t.Fatal(err)
			}
		} else {
			err := cache.Set(fmt.Sprintf("%d", i), i, file.New("none.file"))
			if err != nil {
				t.Fatal(err)
			}
		}
	}

	for i := 0; i < 2000; i++ {
		value, err := cache.Get(fmt.Sprintf("%d", i))
		if err != nil {
			t.Fatal(err)
		}

		if i%2 != 0 && value != nil {
			t.Fatalf("item %d error: it should be expired", i)
		}
	}

	err = cache.Clear()
	if err != nil {
		t.Fatal(err)
	}
}
