package multilevel

import (
	"github.com/landjur/go-caching/container"
)

// New returns a new multilevel caching container.
func New(containers ...container.Container) container.Container {
	return &multilevel{
		containers: containers,
	}
}

type multilevel struct {
	containers []container.Container
}

func (this multilevel) Get(key string) (interface{}, error) {
	for _, v := range this.containers {
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

func (this multilevel) Set(key string, value interface{}) error {
	for _, v := range this.containers {
		err := v.Set(key, value)
		if err != nil {
			return err
		}
	}

	return nil
}

func (this multilevel) Remove(key string) error {
	for _, v := range this.containers {
		err := v.Remove(key)
		if err != nil {
			return err
		}
	}

	return nil
}

func (this multilevel) Clear() error {
	for _, v := range this.containers {
		err := v.Clear()
		if err != nil {
			return err
		}
	}

	return nil
}
