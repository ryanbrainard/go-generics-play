package futures

import (
	"context"
	"github.com/ryanbrainard/go-generics-play/monads/results"
)

type chanFuture[V any] struct {
	done   chan struct{}
	result results.Result[V]
}

func ExecuteChanFuture[V any](ctx context.Context, task func(context.Context) results.Result[V]) Future[V] {
	f := &chanFuture[V]{
		done: make(chan struct{}),
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
