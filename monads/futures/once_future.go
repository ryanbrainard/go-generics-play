package futures

import (
	"context"
	"github.com/ryanbrainard/go-generics-play/monads/results"
	"sync"
)

type onceFuture[V any] struct {
	once     sync.Once
	runnable func()
	result   results.Result[V]
}

func NewOnceFuture[V any](ctx context.Context, task func(context.Context) results.Result[V]) Future[V] {
	f := &onceFuture[V]{}

	f.runnable = func() {
		f.result = task(ctx)
	}

	return f
}

func (f *onceFuture[V]) Run() {
	f.once.Do(f.runnable)
}

func (f *onceFuture[V]) Get() results.Result[V] {
	f.Run()
	return f.result
}
