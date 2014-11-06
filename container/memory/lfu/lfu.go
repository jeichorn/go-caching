package lfu

import (
	"container/heap"
	"github.com/landjur/go-caching/container"
	"github.com/landjur/go-caching/container/memory"
)

// New returns a new in-memory caching container using lfu (least frequently used) arithmetic.
func New(capacity int) container.Container {
	return &lfu{
		capacity: capacity,
		list:     make(items, 0),
		table:    make(map[string]*item),
	}
}

// register the container.
func init() {
	memory.LFU.Register(New)
}

// item represents a caching item.
type item struct {
	Key   string
	Value interface{}
	Index int // index of item in the heap
	Count int // accessed count
}

// items is a collection of item implemented heap.Interface.
type items []*item

func (this items) Len() int {
	return len(this)
}

func (this items) Less(i, j int) bool {
	return this[i].Count < this[j].Count
}

func (this items) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
	this[i].Index = i
	this[j].Index = j
}

func (this *items) Push(x interface{}) {
	index := len(*this)
	item := x.(*item)
	item.Index = index
	*this = append(*this, item)
}

func (this *items) Pop() interface{} {
	old := *this
	index := len(old)
	item := old[index-1]
	item.Index = -1 // for safety
	*this = old[0 : index-1]

	return item
}

type lfu struct {
	capacity int
	list     items
	table    map[string]*item
}

func (this *lfu) Get(key string) (interface{}, error) {
	if len(this.list) == 0 {
		return nil, nil
	}

	if e, ok := this.table[key]; ok {
		e.Count++
		heap.Fix(&this.list, e.Index)
		return e.Value, nil
	}

	return nil, nil
}

func (this *lfu) Set(key string, value interface{}) error {
	if len(this.list) == 0 {
		heap.Init(&this.list)
		this.table = make(map[string]*item)
	}

	if e, ok := this.table[key]; ok {
		e.Value = value
	} else {
		if this.capacity > 0 && len(this.list) == this.capacity {
			item := heap.Pop(&this.list).(*item)
			delete(this.table, item.Key)
		}

		item := &item{
			Key:   key,
			Value: value,
		}
		heap.Push(&this.list, item)
		this.table[key] = item
	}

	return nil
}

func (this *lfu) Remove(key string) error {
	if len(this.list) == 0 {
		return nil
	}

	if e, ok := this.table[key]; ok {
		heap.Remove(&this.list, e.Index)
		delete(this.table, e.Key)
	}

	return nil
}

func (this *lfu) Clear() error {
	if len(this.list) == 0 {
		return nil
	}

	this.list = make(items, 0)
	this.table = make(map[string]*item)
	return nil
}
