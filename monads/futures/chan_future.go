package futures

import (
	"context"
	"github.com/ryanbrainard/go-generics-play/monads/results"
)

type chanFuture[V any] struct {
	cancel context.CancelFunc
	done   chan struct{}
	result results.Result[V]
}

func NewChanFuture[V any](task func(ctx context.Context) results.Result[V]) Future[V] {
	ctx, cancel := context.WithCancel(context.Background())
	f := &chanFuture[V]{
		cancel: cancel,
		done:   make(chan struct{}),
	}

	go func() {
		f.result = task(ctx)
		close(f.done)
	}()

	return f
}

func (f *chanFuture[V]) Get() results.Result[V] {
	<-f.done

	return f.result
}

func (f *chanFuture[V]) Cancel() {
	f.cancel()
}
