package events

import (
	"github.com/google/uuid"
)

type InsertCoinEvent struct {
	id        uuid.UUID
	source    string
	timestamp int64
}

func InsertCoin(id uuid.UUID, source string, ts int64) InsertCoinEvent {
	insert := InsertCoinEvent{
		id:        id,
		source:    source,
		timestamp: ts,
	}

	return insert
}

func (c InsertCoinEvent) ID() uuid.UUID {
	return c.id
}

func (c InsertCoinEvent) Source() string {
	return c.source
}

func (c InsertCoinEvent) Name() string {
	return "Insert Coin"
}

func (c InsertCoinEvent) Timestamp() int64 {
	return c.timestamp
}
