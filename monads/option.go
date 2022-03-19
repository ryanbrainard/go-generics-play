package monads

type Option[V any] interface {
	Get() V
	Map(fn func(V) V) Option[V]
}

type Some[V any] struct {
	v V
}

func (o Some[V]) Get() V {
	return o.v
}

func (o Some[V]) Map(fn func(v V) V) Option[V] {
	return Some[V]{fn(o.v)}
}

type None[V any] struct{}

func (o None[V]) Get() V {
	return *new(V)
}

func (o None[V]) Map(fn func(v V) V) Option[V] {
	return None[V]{}
}
