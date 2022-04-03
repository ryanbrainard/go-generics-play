package futures

import (
	"context"
	"github.com/ryanbrainard/go-generics-play/monads/results"
)

var _ Future[struct{}] = &chanFuture[struct{}]{}
var _ Running = &chanFuture[struct{}]{}

type chanFuture[V any] struct {
	done   chan struct{}
	result results.Result[V]
}

func NewChanFuture[V any](ctx context.Context, task func(context.Context) results.Result[V]) Future[V] {
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

func (f *chanFuture[V]) Running() bool {
	select {
	case <-f.done:
		return false
	default:
		return true
	}
}
