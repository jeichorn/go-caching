package file

import (
	"encoding/gob"
	"github.com/landjur/go-caching/dependency"
	"os"
	"time"
)

// register the type for gob encoding.
func init() {
	gob.Register(time.Time{})
	gob.Register(&file{})
}

// New returns a new file caching dependency.
func New(filePath string) dependency.Dependency {
	var lastModifiedTime time.Time
	fi, err := os.Stat(filePath)
	if err != nil {
		lastModifiedTime = time.Now()
	} else {
		lastModifiedTime = fi.ModTime()
	}

	return &file{
		Path:             filePath,
		LastModifiedTime: lastModifiedTime,
	}
}

type file struct {
	Path             string
	LastModifiedTime time.Time
}

func (this file) HasChanged() bool {
	fi, err := os.Stat(this.Path)
	if err != nil {
		return true
	}

	return fi.ModTime().After(this.LastModifiedTime)
}
