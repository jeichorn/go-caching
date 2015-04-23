package caching

// Dependency represents an external expiration dependency.
type Dependency interface {
	HasChanged() bool
}
