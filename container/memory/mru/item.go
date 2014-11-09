package mru

import (
	"container/list"
)

type item struct {
	Key   string
	Value interface{}
}

func newItems() *items {
	return &items{
		list:  list.New(),
		table: make(map[string]*list.Element),
	}
}

type items struct {
	list  *list.List
	table map[string]*list.Element
}

func (this items) Count() int {
	return this.list.Len()
}

func (this items) Contains(key string) bool {
	_, ok := this.table[key]
	return ok
}

func (this items) Get(key string) interface{} {
	if this.list == nil {
		return nil
	}

	if element, ok := this.table[key]; ok {
		this.list.MoveToFront(element)
		return element.Value.(*item).Value
	}

	return nil
}

func (this *items) Set(key string, value interface{}) {
	if this.list == nil {
		this.list = list.New()
		this.table = make(map[string]*list.Element)
	}

	if element, ok := this.table[key]; ok {
		this.list.MoveToFront(element)
		element.Value.(*item).Value = value
	} else {
		element := this.list.PushFront(&item{Key: key, Value: value})
		this.table[key] = element
	}
}

func (this *items) Discard() (string, interface{}) {
	element := this.list.Front()
	if element == nil {
		return "", nil
	}

	item := element.Value.(*item)
	this.list.Remove(element)
	delete(this.table, item.Key)

	return item.Key, item.Value
}

func (this *items) Remove(key string) {
	if element, ok := this.table[key]; ok {
		this.list.Remove(element)
		delete(this.table, key)
	}
}

func (this *items) Clear() {
	if this.list != nil {
		this.list.Init()
		this.table = make(map[string]*list.Element)
	}
}
