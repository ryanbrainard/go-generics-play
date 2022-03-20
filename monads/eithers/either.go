package eithers

type Either[L, R any] interface {
	// GetOrElse returns the value from this Right or the given argument if this is a Left.
	GetOrElse(R) R

	// Map applies the given function if this is a Right.
	// To return a type other than V, see eithers.Map function.
	Map(fn func(R) R) Either[L, R]

	// OrElse returns this Right or the given argument if this is a Left
	OrElse(Either[L, R]) Either[L, R]

	Split() (L, R)

	Swap() Either[R, L]
}

type right[L, R any] struct {
	r R
}

func Right[L, R any](r R) Either[L, R] {
	return right[L, R]{r}
}

func (e right[L, R]) GetOrElse(R) R {
	return e.r
}

func (e right[L, R]) Map(fn func(R) R) Either[L, R] {
	return Map[L, R, R](e, fn)
}

func (e right[L, R]) OrElse(Either[L, R]) Either[L, R] {
	return e
}

func (e right[L, R]) Split() (L, R) {
	return *new(L), e.r
}

func (e right[L, R]) Swap() Either[R, L] {
	return left[R, L]{e.r}
}

type left[L, R any] struct {
	l L
}

func Left[L, R any](l L) Either[L, R] {
	return left[L, R]{l}
}

func (e left[L, R]) GetOrElse(or R) R {
	return or
}

func (e left[L, R]) Map(fn func(R) R) Either[L, R] {
	return Map[L, R, R](e, fn)
}

func (e left[L, R]) OrElse(or Either[L, R]) Either[L, R] {
	return or
}

func (e left[L, R]) Split() (L, R) {
	return e.l, *new(R)
}

func (e left[L, R]) Swap() Either[R, L] {
	return right[R, L]{e.l}
}

// Map applies the given function if this is a Right.
// To return same type V, see Either.Map method.
func Map[L, R, S any](e Either[L, R], fn func(R) S) Either[L, S] {
	switch e := e.(type) {
	case left[L, R]:
		return left[L, S]{e.l}
	case right[L, R]:
		return right[L, S]{fn(e.r)}
	default:
		panic("impossible")
	}
}

func FlatMap[L, R, S any](e Either[L, R], fn func(R) Either[L, S]) Either[L, S] {
	switch e := e.(type) {
	case left[L, R]:
		return left[L, S]{e.l}
	case right[L, R]:
		return fn(e.r)
	default:
		panic("impossible")
	}
}

// Fold applies fl if this is a Left or fr if this is a Right
func Fold[L, R, C any](e Either[L, R], fl func(L) C, fr func(R) C) C {
	switch e := e.(type) {
	case left[L, R]:
		return fl(e.l)
	case right[L, R]:
		return fr(e.r)
	default:
		panic("impossible")
	}
}
