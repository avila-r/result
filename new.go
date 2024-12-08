package result

func Of[T any](v T, e error) With[T] {
	return With[T]{
		value: &v,
		err:   e,
	}
}

func Ok[T any](value T) With[T] {
	return With[T]{
		value: &value,
		err:   nil,
	}
}

func Error[T any](err error) With[T] {
	return With[T]{
		value: nil,
		err:   err,
	}
}

func Err[T any](err error) With[T] {
	return Error[T](err)
}
