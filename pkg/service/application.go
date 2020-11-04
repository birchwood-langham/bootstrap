package service

import (
	"context"
)

type InitFunc func(ctx context.Context, state StateStore) error
type CleanupFunc func() error

type Application struct {
	initFunctions    []InitFunc
	cleanupFunctions []CleanupFunc
}

func NewApplication() Application {
	return Application{
		initFunctions:    make([]InitFunc, 0),
		cleanupFunctions: make([]CleanupFunc, 0),
	}
}

func (a Application) Init(ctx context.Context, state StateStore) error {
	for _, f := range a.initFunctions {
		if err := f(ctx, state); err != nil {
			return err
		}
	}

	return nil
}

func (a Application) AddInitFunc(initFuncs ...InitFunc) Application {
	a.initFunctions = append(a.initFunctions, initFuncs...)
	return a
}

func (a Application) Cleanup() error {
	for _, f := range a.cleanupFunctions {
		if err := f(); err != nil {
			return err
		}
	}

	return nil
}

func (a Application) AddCleanupFunc(fns ...CleanupFunc) Application {
	a.cleanupFunctions = append(a.cleanupFunctions, fns...)
	return a
}

func (a Application) SetProperties(usage, shortDesc, longDesc string) {
	SetProperties(usage, shortDesc, longDesc)
}
