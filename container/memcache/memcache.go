package memcache

import (
	"bytes"
	"encoding/gob"
	mc "github.com/rainycape/memcache"
	"github.com/wayn3h0/go-caching"
)

// New returns a new instance of caching.Container: memcached caching container.
func New(servers ...string) (caching.Container, error) {
	client, err := mc.New(servers...)
	if err != nil {
		return nil, err
	}

	return &container{
		Client: client,
	}, nil
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

type container struct {
	Client *mc.Client
}

func (this container) Get(key string) (interface{}, error) {
	mci, err := this.Client.Get(key)
	if err != nil {
		return nil, err
	}

	item, err := decode(mci.Value)
	if err != nil {
		return nil, err
	}

	return item.Value, nil
}

func (this container) Set(key string, value interface{}) error {
	data, err := encode(&item{value})
	if err != nil {
		return err
	}

	mci := &mc.Item{
		Key:   key,
		Value: data,
	}

	return this.Client.Set(mci)
}

func (this container) Remove(key string) error {
	err := this.Client.Delete(key)
	if err != mc.ErrCacheMiss {
		return err
	}

	return nil
}

func (this container) Clear() error {
	return this.Client.Flush(0)
}
