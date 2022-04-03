package futures

import (
	"context"
	"github.com/ryanbrainard/go-generics-play/monads/results"
	"sync"
)

type wgFuture[V any] struct {
	wg     sync.WaitGroup
	result results.Result[V]
}

func NewWgFuture[V any](ctx context.Context, task func(context.Context) results.Result[V]) Future[V] {
	f := &wgFuture[V]{}
	f.wg.Add(1)

	go func() {
		defer f.wg.Done()
		f.result = task(ctx)
	}()

	return f
}

func (f *wgFuture[V]) Get() results.Result[V] {
	f.wg.Wait()

	return f.result
}
