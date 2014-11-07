package lru

import (
	"container/list"
)

type entry struct {
	Key   string
	Value interface{}
}

func newContainer() *container {
	return &container{
		List: list.New(),
		Map:  make(map[string]*list.Element),
	}
}

type container struct {
	List *list.List
	Map  map[string]*list.Element
}

func (this container) Count() int {
	return this.List.Len()
}

func (this container) Contains(key string) bool {
	_, ok := this.Map[key]
	return ok
}

func (this container) Get(key string) interface{} {
	if this.List == nil {
		return nil
	}

	if element, ok := this.Map[key]; ok {
		this.List.MoveToFront(element)
		return element.Value.(*entry).Value
	}

	return nil
}

func (this *container) Set(key string, value interface{}) {
	if this.List == nil {
		this.List = list.New()
		this.Map = make(map[string]*list.Element)
	}

	if element, ok := this.Map[key]; ok {
		this.List.MoveToFront(element)
		element.Value.(*entry).Value = value
	} else {
		element := this.List.PushFront(&entry{Key: key, Value: value})
		this.Map[key] = element
	}
}

func (this *container) Discard() (string, interface{}) {
	element := this.List.Back()
	if element == nil {
		return "", nil
	}

	entry := element.Value.(*entry)
	this.List.Remove(element)
	delete(this.Map, entry.Key)

	return entry.Key, entry.Value
}

func (this *container) Remove(key string) {
	if element, ok := this.Map[key]; ok {
		this.List.Remove(element)
		delete(this.Map, key)
	}
}

func (this *container) Clear() {
	this.List.Init()
	this.Map = make(map[string]*list.Element)
}
