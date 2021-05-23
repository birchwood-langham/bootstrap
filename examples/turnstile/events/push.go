package events

import (
	"github.com/google/uuid"
)

type PushEvent struct {
	id        uuid.UUID
	source    string
	timestamp int64
}

func Push(id uuid.UUID, source string, ts int64) PushEvent {
	p := PushEvent{
		id:        id,
		source:    source,
		timestamp: ts,
	}

	return p
}

func (p PushEvent) ID() uuid.UUID {
	return p.id
}

func (p PushEvent) Source() string {
	return p.source
}

func (p PushEvent) Name() string {
	return "Push Turnstile"
}

func (p PushEvent) Timestamp() int64 {
	return p.timestamp
}
