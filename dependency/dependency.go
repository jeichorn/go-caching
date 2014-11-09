package dependency

// Dependency represents the interface of an expiration dependency.
type Dependency interface {
	// HasChanged reports the dependency has changed.
	HasChanged() bool
}
