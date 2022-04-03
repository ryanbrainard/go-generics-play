package results

import "github.com/ryanbrainard/go-generics-play/monads/eithers"

type Result[R any] struct {
	e eithers.Either[error, R]
}

// GetOrElse returns the success value from this Result or the given argument if this is an error.
func (r Result[R]) GetOrElse(or R) R {
	return r.e.GetOrElse(or)
}

// Map returns a new result with function applied on success.
// To return a Result type other than R, see results.Map function.
func (r Result[R]) Map(onSuccess func(R) R) Result[R] {
	return Map(r, onSuccess)
}

// FlatMap returns a new result with function returning a Result applied on success.
// To return a Result type other than R, see results.FlatMap function.
func (r Result[R]) FlatMap(onSuccess func(R) Result[R]) Result[R] {
	return FlatMap(r, onSuccess)
}

// MapError returns a new result with function applied on error.
// Same as results.MapError function.
func (r Result[R]) MapError(onError func(error) error) Result[R] {
	return MapError(r, onError)
}

// Split returns R and error as separate values
func (r Result[R]) Split() (R, error) {
	return r.e.Swap().Split()
}

// New creates a new Result from R and error, choosing the error is non-null
func New[R any](v R, err error) Result[R] {
	if err != nil {
		return Err[R](err)
	} else {
		return Ok[R](v)
	}
}

// Ok creates a new Result from R only
func Ok[R any](v R) Result[R] {
	return Result[R]{eithers.Right[error, R](v)}
}

// Err creates a new Result from error only
func Err[R any](err error) Result[R] {
	return Result[R]{eithers.Left[error, R](err)}
}

// Map returns a new result with function applied on success.
// To return a Result type of R, see Result.Map method.
func Map[R, S any](r Result[R], onSuccess func(R) S) Result[S] {
	return Result[S]{eithers.Map(r.e, onSuccess)}
}

// FlatMap returns a new result with function returning a Result applied on success.
// To return a Result type of R, see Result.FlatMap method.
func FlatMap[R, S any](r Result[R], onSuccess func(R) Result[S]) Result[S] {
	return Result[S]{eithers.FlatMap[error, R, S](r.e, func(r R) eithers.Either[error, S] {
		return onSuccess(r).e
	})}
}

// MapError returns a new result with function applied on error.
// Same as Result.MapError method.
func MapError[R any](r Result[R], onError func(error) error) Result[R] {
	return Result[R]{eithers.Map(r.e.Swap(), onError).Swap()}
}

// Fold returns a new type C on success or error by applying the provided functions.
func Fold[R, C any](r Result[R], onSuccess func(R) C, onError func(error) C) C {
	return eithers.Fold(r.e, onError, onSuccess)
}

// SplitFold returns a new type C on success or error by applying the provided function.
func SplitFold[R, C any](r Result[R], onSuccessOrError func(R, error) C) C {
	return onSuccessOrError(r.Split())
}
