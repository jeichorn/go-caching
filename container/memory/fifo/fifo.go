package fifo

import (
	"container/list"
	"github.com/landjur/go-caching/container"
	"github.com/landjur/go-caching/container/memory"
)

// New returns a new in-memory caching container using fifo (first in first out) arithmetic.
func New(capacity int) container.Container {
	return &fifo{
		capacity: capacity,
		list:     list.New(),
		table:    make(map[string]*list.Element),
	}
}

// register the container.
func init() {
	memory.FIFO.Register(New)
}

type item struct {
	Key   string
	Value interface{}
}

type fifo struct {
	capacity int
	list     *list.List
	table    map[string]*list.Element
}

func (this *fifo) Get(key string) (interface{}, error) {
	if this.list == nil {
		return nil, nil
	}

	if element, ok := this.table[key]; ok {
		return element.Value.(*item).Value, nil
	}

	return nil, nil
}

func (this *fifo) Set(key string, value interface{}) error {
	if this.list == nil {
		this.list = list.New()
		this.table = make(map[string]*list.Element)
	}

	if element, ok := this.table[key]; ok {
		element.Value.(*item).Value = value
	} else {
		if this.capacity > 0 && this.list.Len() == this.capacity {
			element := this.list.Front()
			item := element.Value.(*item)
			this.list.Remove(element)
			delete(this.table, item.Key)
		}

		item := &item{
			Key:   key,
			Value: value,
		}
		element := this.list.PushBack(item)
		this.table[key] = element
	}

	return nil
}

func (this *fifo) Remove(key string) error {
	if this.list == nil {
		return nil
	}

	if element, ok := this.table[key]; ok {
		this.list.Remove(element)
		delete(this.table, key)
	}

	return nil
}

func (this *fifo) Clear() error {
	if this.list == nil {
		return nil
	}

	this.list.Init()
	this.table = make(map[string]*list.Element)
	return nil
}
