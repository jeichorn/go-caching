package lru

import (
	"testing"
)

func Test(t *testing.T) {
	lru := New(3)
	lru.Set("key1", "key1")
	lru.Set("key2", "key2")
	lru.Set("key3", "key3")
	lru.Get("key2")
	lru.Get("key3")
	lru.Set("key4", "key4") // remove key1
	lru.Get("key4")
	lru.Set("key5", "key5") // remove key2

	if item, _ := lru.Get("key3"); item == nil {
		t.Fatal("lru: key3 should exists")
	}
	if item, _ := lru.Get("key4"); item == nil {
		t.Fatal("lru: key4 should exists")
	}
	if item, _ := lru.Get("key5"); item == nil {
		t.Fatal("lru: key5 should exists")
	}
}
