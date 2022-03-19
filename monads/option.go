package monads

type Option[V any] interface {
	Get() V
	GetOrElse(V) V
	Map(fn func(V) V) Option[V]
}

type some[V any] struct {
	v V
}

func SomeOf[V any](v V) some[V] {
	return some[V]{v}
}

func (o some[V]) Get() V {
	return o.v
}

func (o some[V]) GetOrElse(V) V {
	return o.Get()
}

func (o some[V]) Map(fn func(v V) V) Option[V] {
	return some[V]{fn(o.v)}
}

type none[V any] struct{}

func NoneOf[V any]() none[V] {
	return none[V]{}
}

func (o none[V]) Get() V {
	return *new(V)
}

func (o none[V]) GetOrElse(v V) V {
	return v
}

func (o none[V]) Map(fn func(v V) V) Option[V] {
	return none[V]{}
}
