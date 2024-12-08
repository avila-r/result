package result

type With[T any] struct {
    value *T
    err   error
}

func (r With[T]) Value() *T {
	return r.value
}

func (r With[T]) Error() error {
	return r.err
}

func (r With[T]) Err() error {
	return r.Error()
}

func (r With[T]) ErrorMessage() string {
	return r.err.Error()
}
