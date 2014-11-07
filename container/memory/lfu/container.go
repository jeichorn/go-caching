package lfu

import (
	"container/heap"
)

type entry struct {
	Key   string
	Value interface{}
	Index int // index of entry in the heap
	Count int // accessed count
}

type entries []*entry

func (this entries) Len() int {
	return len(this)
}

func (this entries) Less(i, j int) bool {
	return this[i].Count < this[j].Count
}

func (this entries) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
	this[i].Index = i
	this[j].Index = j
}

func (this *entries) Push(x interface{}) {
	index := len(*this)
	entry := x.(*entry)
	entry.Index = index
	*this = append(*this, entry)
}

func (this *entries) Pop() interface{} {
	old := *this
	index := len(old)
	entry := old[index-1]
	entry.Index = -1 // for safety
	*this = old[0 : index-1]

	return entry
}

func newContainer() *container {
	return &container{
		Heap: make(entries, 0),
		Map:  make(map[string]*entry),
	}
}

type container struct {
	Heap entries
	Map  map[string]*entry
}

func (this container) Count() int {
	return this.Heap.Len()
}

func (this container) Contains(key string) bool {
	_, ok := this.Map[key]
	return ok
}

func (this container) Get(key string) interface{} {
	if len(this.Heap) == 0 {
		return nil
	}

	if element, ok := this.Map[key]; ok {
		element.Count++
		heap.Fix(&this.Heap, element.Index)
		return element.Value
	}

	return nil
}

func (this *container) Set(key string, value interface{}) {
	if len(this.Heap) == 0 {
		heap.Init(&this.Heap)
		this.Map = make(map[string]*entry)
	}

	if element, ok := this.Map[key]; ok {
		element.Value = value
	} else {
		entry := &entry{
			Key:   key,
			Value: value,
		}
		heap.Push(&this.Heap, entry)
		this.Map[key] = entry
	}
}

func (this *container) Discard() (string, interface{}) {
	if len(this.Heap) == 0 {
		return "", nil
	}

	entry := heap.Pop(&this.Heap).(*entry)
	delete(this.Map, entry.Key)

	return entry.Key, entry.Value
}

func (this *container) Remove(key string) {
	if element, ok := this.Map[key]; ok {
		heap.Remove(&this.Heap, element.Index)
		delete(this.Map, key)
	}
}

func (this *container) Clear() {
	if len(this.Heap) > 0 {
		this.Heap = make(entries, 0)
		this.Map = make(map[string]*entry)
	}
}
