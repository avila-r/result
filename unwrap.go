package result

func (r With[T]) Get() T {
	return *r.value
}

func (r With[T]) Unwrap() T {
	if r.IsError() {
		panic(r.err)
	}

    return *r.value
}

func (r With[T]) Expect(message ...string) T {
	var msg string
	if len(message) > 0 {
		msg = message[0]
	} else {
		msg = r.err.Error()
	}

	if r.IsError() {
		panic(msg)
	}

    return *r.value
}
