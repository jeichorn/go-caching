package multilevel

import (
	"github.com/wayn3h0/go-caching"
)

// New returns a new instance of caching.Container: container caching container wrapper for container caching.
func New(containers ...caching.Container) caching.Container {
	return &container{
		list: containers,
	}
}

type container struct {
	list []caching.Container
}

func (this container) Get(key string) (interface{}, error) {
	for _, v := range this.list {
		value, err := v.Get(key)
		if err != nil {
			return nil, err
		}

		if value != nil {
			return value, nil
		}
	}

	return nil, nil
}

func (this container) Set(key string, value interface{}) error {
	for _, v := range this.list {
		err := v.Set(key, value)
		if err != nil {
			return err
		}
	}

	return nil
}

func (this container) Remove(key string) error {
	for _, v := range this.list {
		err := v.Remove(key)
		if err != nil {
			return err
		}
	}

	return nil
}

func (this container) Clear() error {
	for _, v := range this.list {
		err := v.Clear()
		if err != nil {
			return err
		}
	}

	return nil
}
