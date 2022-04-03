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

func NewWgFuture[V any](task func(ctx context.Context) results.Result[V]) Future[V] {
	f := &wgFuture[V]{}
	f.wg.Add(1)
	go func() {
		defer f.wg.Done()
		f.result = task(context.TODO())
	}()
	return f
}

func (f *wgFuture[V]) Get() results.Result[V] {
	f.wg.Wait()

	return f.result
}
