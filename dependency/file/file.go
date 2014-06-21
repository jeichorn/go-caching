package file

import (
	"github.com/landjur/go-caching"
	"os"
	"time"
)

// New returns a new instance of File.
func New(filename string) *File {
	var lastModifiedTime time.Time
	fi, err := os.Stat(filename)
	if err != nil {
		lastModifiedTime = time.Now()
	} else {
		lastModifiedTime = fi.ModTime()
	}

	return &File{
		Name:             filename,
		lastModifiedTime: lastModifiedTime,
	}
}

// File represents a caching dependency policy by file.
type File struct {
	Name             string
	lastModifiedTime time.Time
}

func (this File) HasExpired(item *caching.Item) bool {
	fi, err := os.Stat(this.Name)
	if err != nil {
		return true
	}

	return fi.ModTime().After(this.lastModifiedTime)
}
