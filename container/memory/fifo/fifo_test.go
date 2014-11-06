package fifo

import (
	"testing"
)

func Test(t *testing.T) {
	fifo := New(3)
	fifo.Set("key1", "key1")
	fifo.Set("key2", "key2")
	fifo.Set("key3", "key3")
	fifo.Set("key4", "key4")
	fifo.Set("key5", "key5")
	if item, _ := fifo.Get("key3"); item == nil {
		t.Fatal("fifo: key3 should exists")
	}
	if item, _ := fifo.Get("key4"); item == nil {
		t.Fatal("fifo: key4 should exists")
	}
	if item, _ := fifo.Get("key5"); item == nil {
		t.Fatal("fifo: key5 should exists")
	}
}
