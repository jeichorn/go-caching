package memcache

import (
	"bytes"
	"encoding/gob"
	mc "github.com/bradfitz/gomemcache/memcache"
	"github.com/landjur/go-caching/container"
)

// New returns a new memcached caching container.
func New(servers ...string) container.Container {
	return &memcache{
		client: mc.New(servers...),
	}
}

// item represents a caching item.
type item struct {
	Value interface{}
}

// encode encodes the item to binary data by gob.
func encode(item *item) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(*item)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

// decode decodes the item from binary data by gob.
func decode(data []byte) (*item, error) {
	decoder := gob.NewDecoder(bytes.NewBuffer(data))
	var item item
	err := decoder.Decode(&item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

type memcache struct {
	client *mc.Client
}

func (this memcache) Get(key string) (interface{}, error) {
	mci, err := this.client.Get(key)
	if err != nil {
		return nil, err
	}

	item, err := decode(mci.Value)
	if err != nil {
		return nil, err
	}

	return item.Value, nil
}

func (this memcache) Set(key string, value interface{}) error {
	data, err := encode(&item{value})
	if err != nil {
		return err
	}

	mci := &mc.Item{
		Key:   key,
		Value: data,
	}

	return this.client.Set(mci)
}

func (this memcache) Remove(key string) error {
	err := this.client.Delete(key)
	if err != mc.ErrCacheMiss {
		return err
	}

	return nil
}

func (this memcache) Clear() error {
	return this.client.DeleteAll()
}
