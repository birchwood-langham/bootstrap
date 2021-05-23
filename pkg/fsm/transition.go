package fsm

// CheckFn takes a state and determines whether or not it is ready to transition
type CheckFn func(State) bool

// NextFn is a function that takes a state and generates the next state
// For example, any properties that should be copied to the new state
// or calculations that need to be performed using the current state's
// data and passed to the next state
type NextFn func(State) State

// Transition contains the check functions required for a transition and the
// Next function to generate the next state
type Transition struct {
	// Checks are the checks that need to pass before the transition
	// can take place
	Checks []CheckFn
	// Next creates the next state for the state machine to transition to
	Next NextFn
}
