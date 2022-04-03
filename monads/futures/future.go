package futures

import (
	"context"
	"github.com/ryanbrainard/go-generics-play/monads/results"
)

type Executor[V any] func(context.Context, func(context.Context) results.Result[V]) Future[V]

type Future[V any] interface {
	Get() results.Result[V]
}
