package result

func Map[T any, U any](r With[T], f func(T) U) With[U] {
	if r.IsError() || r.IsEmpty() {
		return Error[U](r.err)
	}

	return Ok(f(*r.value))
}

func UnwrapOrElse[T any](r With[T], fallback func() T) T {
	if r.IsError() || r.IsEmpty() {
		return fallback()
	}

	return *r.value
}

func Flatten[T any](r With[With[T]]) With[T] {
	if r.IsError() || r.IsEmpty() {
		return Error[T](r.err)
	}

	return *r.value
}

func Combine[T any, U any](w1 With[T], w2 With[U]) With[struct {
	First  T
	Second U
}] {
	if w1.IsError() || w1.IsEmpty() {
		return Error[struct {
			First  T
			Second U
		}](w1.err)
	}

	if w2.IsError() || w2.IsEmpty() {
		return Error[struct {
			First  T
			Second U
		}](w2.err)
	}

	return Ok(struct {
		First  T
		Second U
	}{First: *w1.value, Second: *w2.value})
}
