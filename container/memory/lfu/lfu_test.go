package lfu

import (
	"testing"
)

func Test(t *testing.T) {
	lfu := New(3)
	lfu.Set("key1", "key1")
	lfu.Set("key2", "key2")
	lfu.Set("key3", "key3")
	lfu.Get("key2")
	lfu.Get("key2")         // key2 count: 2
	lfu.Get("key3")         // key3 count: 1
	lfu.Set("key4", "key4") // remove key1
	lfu.Get("key4")
	lfu.Get("key4")         // key4 count: 2
	lfu.Set("key5", "key5") // remove key3

	if item, _ := lfu.Get("key2"); item == nil {
		t.Fatal("lfu: key2 should exists")
	}
	if item, _ := lfu.Get("key4"); item == nil {
		t.Fatal("lfu: key4 should exists")
	}
	if item, _ := lfu.Get("key5"); item == nil {
		t.Fatal("lfu: key5 should exists")
	}
}
