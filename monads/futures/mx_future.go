package futures

import (
	"context"
	"github.com/ryanbrainard/go-generics-play/monads/results"
	"sync"
)

type mxFuture[V any] struct {
	mx     sync.Mutex
	result results.Result[V]
}

func NewMxFuture[V any](ctx context.Context, task func(context.Context) results.Result[V]) Future[V] {
	f := &mxFuture[V]{}
	f.mx.Lock()

	go func() {
		defer f.mx.Unlock()
		f.result = task(ctx)
	}()

	return f
}

func (f *mxFuture[V]) Get() results.Result[V] {
	f.mx.Lock()
	defer f.mx.Unlock()

	return f.result
}

func (f *mxFuture[V]) Running() bool {
	locked := f.mx.TryLock()
	if locked {
		f.mx.Unlock()
	}
	return !locked
}
