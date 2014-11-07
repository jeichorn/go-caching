package lfu

import (
	"container/heap"
)

type item struct {
	Key   string
	Value interface{}
	Index int // index of item in the heap
	Count int // accessed count
}

type itemsHeap []*item

func (this itemsHeap) Len() int {
	return len(this)
}

func (this itemsHeap) Less(i, j int) bool {
	return this[i].Count < this[j].Count
}

func (this itemsHeap) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
	this[i].Index = i
	this[j].Index = j
}

func (this *itemsHeap) Push(x interface{}) {
	index := len(*this)
	item := x.(*item)
	item.Index = index
	*this = append(*this, item)
}

func (this *itemsHeap) Pop() interface{} {
	old := *this
	index := len(old)
	item := old[index-1]
	item.Index = -1 // for safety
	*this = old[0 : index-1]

	return item
}

func newItems() *items {
	return &items{
		list:  make(itemsHeap, 0),
		table: make(map[string]*item),
	}
}

type items struct {
	list  itemsHeap
	table map[string]*item
}

func (this items) Count() int {
	return this.list.Len()
}

func (this items) Contains(key string) bool {
	_, ok := this.table[key]
	return ok
}

func (this items) Get(key string) interface{} {
	if len(this.list) == 0 {
		return nil
	}

	if element, ok := this.table[key]; ok {
		element.Count++
		heap.Fix(&this.list, element.Index)
		return element.Value
	}

	return nil
}

func (this *items) Set(key string, value interface{}) {
	if len(this.list) == 0 {
		heap.Init(&this.list)
		this.table = make(map[string]*item)
	}

	if element, ok := this.table[key]; ok {
		element.Value = value
	} else {
		item := &item{
			Key:   key,
			Value: value,
		}
		heap.Push(&this.list, item)
		this.table[key] = item
	}
}

func (this *items) Discard() (string, interface{}) {
	if len(this.list) == 0 {
		return "", nil
	}

	item := heap.Pop(&this.list).(*item)
	delete(this.table, item.Key)

	return item.Key, item.Value
}

func (this *items) Remove(key string) {
	if element, ok := this.table[key]; ok {
		heap.Remove(&this.list, element.Index)
		delete(this.table, key)
	}
}

func (this *items) Clear() {
	if len(this.list) > 0 {
		this.list = make(itemsHeap, 0)
		this.table = make(map[string]*item)
	}
}
