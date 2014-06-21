package caching

// Priority represents priority of cache item.
type Priority byte

const (
	PriorityLow Priority = iota
	PriorityNormal
	PriorityHigh
	PriorityNotRemovable
)
