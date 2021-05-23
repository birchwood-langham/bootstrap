package states

import (
	"errors"

	"github.com/birchwood-langham/bootstrap/examples/turnstile/events"
	"github.com/birchwood-langham/bootstrap/pkg/fsm"
	"github.com/google/uuid"
)

type LockedState struct {
	id          uuid.UUID
	HasCoin     bool
	transitions []fsm.Transition
}

func Locked(id uuid.UUID) *LockedState {
	locked := LockedState{
		id:      id,
		HasCoin: false,
	}

	return &locked
}

func (l *LockedState) ID() uuid.UUID {
	return l.id
}

func (l *LockedState) Description() string {
	return "Turnstile is locked"
}

func (l *LockedState) Execute(evt fsm.Event) error {
	switch e := evt.(type) {
	case events.InsertCoinEvent:
		l.HasCoin = true
	case events.PushEvent:
		return errors.New("you have to insert a coin first")
	default:
		return fsm.UnexpectedEventError(e)
	}

	return nil
}

func (l *LockedState) Next() fsm.State {
	return fsm.Next(l, l.transitions...)
}

func (l *LockedState) WithTransitions(transitions ...fsm.Transition) fsm.State {
	l.transitions = append(l.transitions, transitions...)
	return l
}
