package futures

import (
	"context"
	"github.com/ryanbrainard/go-generics-play/monads/results"
)

type chanFuture[V any] struct {
	done   chan struct{}
	result results.Result[V]
}

func NewChanFuture[V any](task func(ctx context.Context) results.Result[V]) Future[V] {
	f := &chanFuture[V]{
		done: make(chan struct{}),
	}
	go func() {
		f.result = task(context.TODO())
		close(f.done)
	}()
	return f
}

func (f *chanFuture[V]) Get() results.Result[V] {
	<-f.done
	return f.result
}
