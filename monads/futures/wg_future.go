package futures

import (
	"context"
	"github.com/ryanbrainard/go-generics-play/monads/results"
	"sync"
)

type wgFuture[V any] struct {
	cancel context.CancelFunc
	wg     sync.WaitGroup
	result results.Result[V]
}

func NewWgFuture[V any](task func(ctx context.Context) results.Result[V]) Future[V] {
	ctx, cancel := context.WithCancel(context.Background())
	f := &wgFuture[V]{cancel: cancel}
	f.wg.Add(1)

	go func() {
		defer f.wg.Done()
		defer f.Cancel()
		f.result = task(ctx)
	}()

	return f
}

func (f *wgFuture[V]) Get() results.Result[V] {
	f.wg.Wait()

	return f.result
}

func (f *wgFuture[V]) Cancel() {
	f.cancel()
}
