package x

func Empty[T any]() T {
	var zero T
	return zero
}

func IsEmpty[T comparable](v T) bool {
	var zero T
	return zero == v
}

func ToPtr[T any](x T) *T {
	return &x
}

func FromPtr[T any](x *T) T {
	if x == nil {
		return Empty[T]()
	}
	return *x
}

func Coalesce[T comparable](v ...T) (result T, ok bool) {
	for _, e := range v {
		if e != result {
			result = e
			ok = true
			return
		}
	}
	return
}
