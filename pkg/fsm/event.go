package fsm

import "github.com/google/uuid"

type Event interface {
	// ID is a unique identifier for the event
	ID() uuid.UUID
	// Source is a unique identifier used to determine where the event came from
	Source() string
	// Name of the event
	Name() string
	// Timestamp is the time of the event as nanoseconds past epoch
	Timestamp() int64
}
