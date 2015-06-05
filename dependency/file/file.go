package file

import (
	"encoding/gob"
	"github.com/jeichorn/go-caching"
	"os"
	"time"
)

// register the type for gob encoding.
func init() {
	gob.Register(new(dependency))
}

// New returns a new instance of caching.Dependency.
func New(filePath string) caching.Dependency {
	var lastModifiedTime time.Time
	fi, err := os.Stat(filePath)
	if err != nil {
		lastModifiedTime = time.Now()
	} else {
		lastModifiedTime = fi.ModTime()
	}

	return &dependency{
		FilePath:         filePath,
		LastModifiedTime: lastModifiedTime,
	}
}

type dependency struct {
	FilePath         string
	LastModifiedTime time.Time
}

func (this dependency) HasChanged() bool {
	fi, err := os.Stat(this.FilePath)
	if err != nil {
		return true
	}

	return fi.ModTime().After(this.LastModifiedTime)
}
