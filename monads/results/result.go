package results

import "go.ryanbrainard.com/go-generics-play/monads/eithers"

type Result[V any] struct {
	e eithers.Either[error, V]
}

// GetOrElse returns the success value from this Result or the given argument if this is an error.
func (r Result[V]) GetOrElse(or V) V {
	return r.e.GetOrElse(or)
}

// Map returns a new result with function applied on success.
// To return a Result type other than V, see results.Map function.
func (r Result[V]) Map(onSuccess func(V) V) Result[V] {
	return Map(r, onSuccess)
}

// MapError returns a new result with function applied on error.
// Same as results.MapError function.
func (r Result[V]) MapError(onError func(error) error) Result[V] {
	return MapError(r, onError)
}

// Split returns V and error as separate values
func (r Result[V]) Split() (V, error) {
	return r.e.Swap().Split()
}

// New creates a new Result from V and error, choosing the error is non-null
func New[V any](v V, err error) Result[V] {
	if err != nil {
		return Err[V](err)
	} else {
		return Ok[V](v)
	}
}

// Ok creates a new Result from V only
func Ok[V any](v V) Result[V] {
	return Result[V]{eithers.Right[error, V](v)}
}

// Err creates a new Result from error only
func Err[V any](err error) Result[V] {
	return Result[V]{eithers.Left[error, V](err)}
}

// Map returns a new result with function applied on success.
// To return a Result type of V, see Result.Map method.
func Map[V, W any](r Result[V], onSuccess func(V) W) Result[W] {
	return Result[W]{eithers.Map(r.e, onSuccess)}
}

// MapError returns a new result with function applied on error.
// Same as Result.MapError method.
func MapError[V any](r Result[V], onError func(error) error) Result[V] {
	return Result[V]{eithers.Map(r.e.Swap(), onError).Swap()}
}

// Fold returns a new type C on success or error by applying the provided functions.
func Fold[V, C any](r Result[V], onSuccess func(V) C, onError func(error) C) C {
	return eithers.Fold(r.e, onError, onSuccess)
}
