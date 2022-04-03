package futures

import (
	"context"
	"github.com/ryanbrainard/go-generics-play/monads/results"
	"sync"
)

type mxFuture[V any] struct {
	cancel context.CancelFunc
	mx     sync.Mutex
	result results.Result[V]
}

func NewMxFuture[V any](task func(ctx context.Context) results.Result[V]) Future[V] {
	ctx, cancel := context.WithCancel(context.Background())
	f := &mxFuture[V]{cancel: cancel}
	f.mx.Lock()

	go func() {
		defer f.mx.Unlock()
		defer f.cancel()
		f.result = task(ctx)
	}()

	return f
}

func (f *mxFuture[V]) Get() results.Result[V] {
	f.mx.Lock()
	defer f.mx.Unlock()

	return f.result
}

func (f *mxFuture[V]) Cancel() {
	f.cancel()
}
