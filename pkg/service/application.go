package service

import (
	"context"
)

type InitFunc func(ctx context.Context, state StateStore) error
type CleanupFunc func(state StateStore) error
type RunFunc func(ctx context.Context, state StateStore) error

type Application struct {
	initFunctions    []InitFunc
	cleanupFunctions []CleanupFunc
	runFunc          RunFunc
}

func NewApplication() Application {
	return Application{
		initFunctions:    make([]InitFunc, 0),
		cleanupFunctions: make([]CleanupFunc, 0),
		runFunc:          nil,
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

func (a Application) Cleanup(state StateStore) error {
	for _, f := range a.cleanupFunctions {
		if err := f(state); err != nil {
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
	SetCliProperties(usage, shortDesc, longDesc)
}

func (a Application) WithRunFunc(f RunFunc) Application {
	a.runFunc = f
	return a
}

func (a Application) RunFunction() RunFunc {
	return a.runFunc
}
