package ptrs

func Of[E any](value E) *E {
	return &value
}

func Deref[E any](ptr *E) (E, bool) {
	if ptr != nil {
		return *ptr, true
	}
	var empty E
	return empty, false
}

func DerefOr[E any](ptr *E, or E) E {
	if ptr != nil {
		return *ptr
	}
	return or
}

func New[E any]() *E {
	return new(E)
}

func Flatten[E any](ptrs []*E) []E {
	var n = 0
	for _, ptr := range ptrs {
		if ptr != nil {
			n++
		}
	}

	var values = make([]E, 0, n)
	for _, ptr := range ptrs {
		if ptr != nil {
			values = append(values, *ptr)
		}
	}
	return values
}

func Ref[E any](values []E) []*E {
	var ptrs = make([]*E, len(values))
	for i := range values {
		ptrs[i] = &values[i]
	}
	return ptrs
}

func Make[E any](n int, fn func(i int, ptr *E)) []*E {
	var values = make([]E, n)
	if fn == nil {
		return Ref(values)
	}

	var ptrs = make([]*E, len(values))
	for i := range values {
		var ptr = &values[i]
		fn(i, ptr)
		ptrs[i] = ptr
	}
	return ptrs
}

func EqualSlice[E comparable, S ~[]*E](a, b S) bool {
	if len(a) != len(b) {
		return false
	}
	for i, pa := range a {
		if !Equal(pa, b[i]) {
			return false
		}
	}
	return true
}

func Equal[E comparable](a, b *E) bool {
	switch {
	case a == b:
		return true
	case a == nil:
		return b == nil
	case b == nil:
		return a == nil
	default:
		return *a == *b
	}
}
