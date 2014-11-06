package lru

import (
	"container/list"
	"github.com/landjur/go-caching/container"
	"github.com/landjur/go-caching/container/memory"
)

// New returns a new in-memory caching container using LRU (least recently used) arithmetic.
func New(capacity int) container.Container {
	return &lru{
		capacity: capacity,
		list:     list.New(),
		table:    make(map[string]*list.Element),
	}
}

// register the container.
func init() {
	memory.LRU.Register(New)
}

type item struct {
	Key   string
	Value interface{}
}

type lru struct {
	capacity int
	list     *list.List
	table    map[string]*list.Element
}

func (this *lru) Get(key string) (interface{}, error) {
	if this.list == nil {
		return nil, nil
	}

	if element, ok := this.table[key]; ok {
		this.list.MoveToFront(element)
		return element.Value.(*item).Value, nil
	}

	return nil, nil
}

func (this *lru) Set(key string, value interface{}) error {
	if this.list == nil {
		this.list = list.New()
		this.table = make(map[string]*list.Element)
	}

	if element, ok := this.table[key]; ok {
		this.list.MoveToFront(element)
		element.Value.(*item).Value = value
	} else {
		if this.capacity > 0 && this.list.Len() == this.capacity {
			element := this.list.Back()
			item := element.Value.(*item)
			this.list.Remove(element)
			delete(this.table, item.Key)
		}

		item := &item{
			Key:   key,
			Value: value,
		}
		element := this.list.PushFront(item)
		this.table[key] = element
	}

	return nil
}

func (this *lru) Remove(key string) error {
	if this.list == nil {
		return nil
	}

	if element, ok := this.table[key]; ok {
		this.list.Remove(element)
		delete(this.table, key)
	}

	return nil
}

func (this *lru) Clear() error {
	if this.list == nil {
		return nil
	}

	this.list.Init()
	this.table = make(map[string]*list.Element)
	return nil
}
