package futures

import (
	"github.com/ryanbrainard/go-generics-play/monads/results"
)

type Future[V any] interface {
	Cancel()
	Get() results.Result[V]
}
