package options

type Option[V any] interface {
	// Fold returns the result of applying f to this Option's value if the Option is a Some.
	// To return a type other than V, see options.Fold function.
	Fold(ifEmpty V) func(func(V) V) V

	// Get returns the Option's value
	Get() V

	// GetOrElse returns the option's value if the Option is a Some; otherwise, return the result of evaluating default.
	GetOrElse(V) V

	// Map returns a Some containing the result of applying fn to this Option's value if this Option is a Some; otherwise None.
	// To return a type other than V, see options.Map function.
	Map(fn func(V) V) Option[V]

	// OrElse returns this Option if it is a Some; otherwise, returns the alternative Option.
	OrElse(Option[V]) Option[V]
}

type some[V any] struct {
	v V
}

func Some[V any](v V) Option[V] {
	return some[V]{v}
}

func (o some[V]) Fold(ifEmpty V) func(func(V) V) V {
	return Fold[V, V](o, ifEmpty)
}

func (o some[V]) Get() V {
	return o.v
}

func (o some[V]) GetOrElse(_ V) V {
	return o.Get()
}

func (o some[V]) Map(fn func(v V) V) Option[V] {
	return Map[V, V](o, fn)
}

func (o some[V]) OrElse(_ Option[V]) Option[V] {
	return o
}

type none[V any] struct{}

func None[V any]() Option[V] {
	return none[V]{}
}

func (o none[V]) Fold(ifEmpty V) func(func(V) V) V {
	return Fold[V, V](o, ifEmpty)
}

func (o none[V]) Get() V {
	return *new(V)
}

func (o none[V]) GetOrElse(v V) V {
	return v
}

func (o none[V]) Map(fn func(v V) V) Option[V] {
	return Map[V, V](o, fn)
}

func (o none[V]) OrElse(b Option[V]) Option[V] {
	return b
}

// Map returns a Some containing the result of applying fn to this Option's value if this Option is a Some; otherwise None.
// To return same type V, see Option.Map method.
func Map[V, R any](o Option[V], fn func(V) R) Option[R] {
	switch o.(type) {
	case some[V]:
		return some[R]{fn(o.Get())}
	default:
		return none[R]{}
	}
}

// Fold returns the result of applying f to this Option's value if the Option is a Some.
// To return same type V, see Option.Fold method.
func Fold[V, R any](o Option[V], ifEmpty R) func(func(V) R) R {
	return func(fn func(V) R) R {
		switch o.(type) {
		case some[V]:
			return fn(o.Get())
		default:
			return ifEmpty
		}
	}
}
