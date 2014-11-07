package fifo

import (
	"container/list"
)

type entry struct {
	Key   string
	Value interface{}
}

func newData() *data {
	return &data{
		List: list.New(),
		Map:  make(map[string]*list.Element),
	}
}

type data struct {
	List *list.List
	Map  map[string]*list.Element
}

func (this data) Count() int {
	return this.List.Len()
}

func (this data) Contains(key string) bool {
	_, ok := this.Map[key]
	return ok
}

func (this data) Get(key string) interface{} {
	if this.List == nil {
		return nil
	}

	if element, ok := this.Map[key]; ok {
		return element.Value.(*entry).Value
	}

	return nil
}

func (this *data) Set(key string, value interface{}) {
	if this.List == nil {
		this.List = list.New()
		this.Map = make(map[string]*list.Element)
	}

	if element, ok := this.Map[key]; ok {
		element.Value.(*entry).Value = value
	} else {
		element := this.List.PushBack(&entry{Key: key, Value: value})
		this.Map[key] = element
	}
}

func (this *data) Discard() (string, interface{}) {
	element := this.List.Front()
	if element == nil {
		return "", nil
	}

	entry := element.Value.(*entry)
	this.List.Remove(element)
	delete(this.Map, entry.Key)

	return entry.Key, entry.Value
}

func (this *data) Remove(key string) {
	if element, ok := this.Map[key]; ok {
		this.List.Remove(element)
		delete(this.Map, key)
	}
}

func (this *data) Clear() {
	this.List.Init()
	this.Map = make(map[string]*list.Element)
}
