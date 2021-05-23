package fsm

import (
	"fmt"

	"github.com/google/uuid"
)

type State interface {
	// ID is the unique identifier of the state so different states of the same nature
	// can be distinguishable
	ID() uuid.UUID
	// Description of the state
	Description() string
	// Execute processes the event that is passed to is
	Execute(Event) error
	// Next checks the current state and determines what state it should transition to
	Next() State
	// WithTransitions sets the transitions that are supported by each state
	WithTransitions(...Transition) State
}

// Next takes the current state and a list of transitions then evaluates each
// transition to see which state it should transition to.
// If no transition check passes, Next will return the current state
func Next(current State, transitions ...Transition) State {
	for _, tr := range transitions {
		failed := false
		// evaluate all check functions
		// if any fail, then we should check the next transition
		for _, ck := range tr.Checks {
			if !ck(current) {
				failed = true
				break
			}
		}

		if !failed {
			// if all checks pass, then we should transition to the next state
			return tr.Next(current)
		}
	}

	// if we are here, then at least one check on each transition must have failed
	// so return the current state
	return current
}

func UnexpectedEventError(event Event) error {
	return fmt.Errorf("an unexpected event %s was received from %s", event.Name(), event.Source())
}
