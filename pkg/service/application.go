package service

import "context"

type InitFunc func(context.Context, Application) error
type CleanupFunc func(Application) error

type Application interface {
	Init(context.Context) error
	AddInitFunc(...InitFunc)
	Cleanup() error
	AddCleanupFunc(...CleanupFunc)
	SetProperties(usage, shortDesc, longDesc string)
}
