package main

import (
	"context"
	"time"

	"go.uber.org/zap"

	"github.com/birchwood-langham/bootstrap/v1/pkg/cmd"
	"github.com/birchwood-langham/bootstrap/v1/pkg/logger"
	"github.com/birchwood-langham/bootstrap/v1/pkg/service"
)

type serverState struct {
	service.StateStore
	notifyCh chan struct{}
	cancel   context.CancelFunc
}

var state *serverState

func initApp(ctx context.Context, store service.StateStore) error {
	log := logger.New(logger.ApplicationLogLevel(), logger.ConfiguredLumberjackLogger())

	ss := store.(*serverState)

	log.Info("initializing application")

	childCtx, cancel := context.WithCancel(ctx)
	notifyCh := make(chan struct{})

	go func(c context.Context, notifyCh chan struct{}) {

		log.Info("Starting long running process...")

		count := 0

		// long running process here
		for {
			select {
			case <-c.Done():
				log.Info("Stopping long running process")
				notifyCh <- struct{}{}

				return
			case <-time.After(5 * time.Second):
				count++
				log.Info("Current count", zap.Int("count", count))
			}
		}
	}(childCtx, notifyCh)

	ss.cancel = cancel
	ss.notifyCh = notifyCh

	return nil
}

func cleanupApp(state service.StateStore) error {
	ss := state.(*serverState)

	ss.cancel()

	// we wait for notification that our long running process has cleanly exited
	<-ss.notifyCh

	return nil
}

func main() {
	app := service.NewApplication().
		AddInitFunc(initApp).
		AddCleanupFunc(cleanupApp)

	app.SetProperties("usage message", "short description", "long description for the application to be displayed when run with the help flag")

	state = &serverState{}

	// NewStateStore() implements the StateStore interface and creates a non-persistent thread safe in-memory state store
	cmd.Execute(context.Background(), app, state)
}
