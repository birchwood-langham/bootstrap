package main

import (
	"context"
	"time"

	"github.com/birchwood-langham/bootstrap/examples/turnstile/events"
	"github.com/birchwood-langham/bootstrap/examples/turnstile/states"
	"github.com/birchwood-langham/bootstrap/examples/turnstile/transitions"
	"github.com/birchwood-langham/bootstrap/pkg/cmd"
	"github.com/birchwood-langham/bootstrap/pkg/fsm"
	"github.com/birchwood-langham/bootstrap/pkg/logger"
	"github.com/birchwood-langham/bootstrap/pkg/service"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const (
	machineName = "Turnstile Service"
)

type TurnstileState struct {
	service.StateStore
	machine            fsm.Machine
	errCh              chan error
	incoming           chan fsm.Event
	cancelMachine      context.CancelFunc
	cancelErrorHandler context.CancelFunc
}

func initStateMachine(ctx context.Context, state service.StateStore) error {
	s := state.(*TurnstileState)
	s.incoming = make(chan fsm.Event)
	s.machine, s.errCh = fsm.New(uuid.New(), machineName, states.Locked(uuid.New()).
		WithTransitions(
			fsm.Transition{
				Checks: []fsm.CheckFn{transitions.HasCoin},
				Next:   transitions.ToUnlocked,
			},
		))

	machineCtx, cancel := context.WithCancel(ctx)

	s.cancelMachine = cancel

	go func(c context.Context, events <-chan fsm.Event) {
		logger.Logger().Info("Starting Turnstile State Machine")
		if err := s.machine.Run(c, events); err != nil {
			logger.Logger().Error("state machine error", zap.Error(err))
			s.cancelMachine()
		}
	}(machineCtx, s.incoming)

	return nil
}

func initErrorHandler(ctx context.Context, state service.StateStore) error {
	s := state.(*TurnstileState)

	errCtx, cancel := context.WithCancel(ctx)

	s.cancelErrorHandler = cancel

	go func(c context.Context, errCh <-chan error) {
		for {
			select {
			case <-ctx.Done():
				logger.Logger().Warn("Stopping turnstile service error handler")
				return
			case err := <-errCh:
				logger.Logger().Error(err.Error())
			}
		}
	}(errCtx, s.errCh)

	return nil
}

func runTurnstile(_ context.Context, state service.StateStore) error {
	s := state.(*TurnstileState)
	s.incoming <- events.Push(uuid.New(), "TEST", time.Now().UnixNano())
	s.incoming <- events.InsertCoin(uuid.New(), "TEST", time.Now().UnixNano())
	s.incoming <- events.InsertCoin(uuid.New(), "TEST", time.Now().UnixNano())
	s.incoming <- events.Push(uuid.New(), "TEST", time.Now().UnixNano())
	s.incoming <- events.Push(uuid.New(), "TEST", time.Now().UnixNano())
	return nil
}

func cleanupStateMachine(state service.StateStore) error {
	s := state.(*TurnstileState)
	s.cancelMachine()
	return nil
}

func cleanupCleanupErrorHandler(state service.StateStore) error {
	s := state.(*TurnstileState)
	s.cancelErrorHandler()
	return nil
}

func main() {
	app := service.NewApplication().
		AddInitFunc(initStateMachine).
		AddInitFunc(initErrorHandler).
		AddCleanupFunc(cleanupStateMachine).
		AddCleanupFunc(cleanupCleanupErrorHandler).
		WithRunFunc(runTurnstile)

	app.SetProperties("", "", "")

	state := new(TurnstileState)

	cmd.Execute(context.Background(), app, state)
}
