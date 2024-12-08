package result

func (r With[T]) IsEmpty() bool {
	return r.value != nil
}

func (r With[T]) IsSuccess() bool {
	return !r.IsError()
}

func (r With[T]) IsError() bool {
	return r.err != nil
}
