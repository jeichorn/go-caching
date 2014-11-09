package mru

import (
	"testing"
)

func Test(t *testing.T) {
	mru := New(3)
	mru.Set("key1", "key1")
	mru.Set("key2", "key2")
	mru.Set("key3", "key3")
	mru.Get("key2")
	mru.Get("key3")
	mru.Set("key4", "key4") // remove key3
	mru.Get("key4")
	mru.Set("key5", "key5") // remove key4

	if item, _ := mru.Get("key1"); item == nil {
		t.Fatal("mru: key1 should exists")
	}
	if item, _ := mru.Get("key2"); item == nil {
		t.Fatal("mru: key2 should exists")
	}
	if item, _ := mru.Get("key5"); item == nil {
		t.Fatal("mru: key5 should exists")
	}
}
