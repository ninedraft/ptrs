package ptrs

// Of returns a pointer to copy of passed value.
//	Of(true)  -> &true
//	Of(false) -> &false
//	Of("hello, world") -> &"hello, world"
func Of[E any](value E) *E {
	return &value
}

// Deref tries to take value from provided pointer.
// Returns falkse, if pointer is nil.
//	Deref(Of(1)) -> (1, true)
//	Deref(nil)   -> (0, false)
func Deref[E any](ptr *E) (E, bool) {
	if ptr != nil {
		return *ptr, true
	}
	var empty E
	return empty, false
}

// DerefOr tries to take value from provided pointer.
// Returns 'or value, if ptr is nil.
//	DerefOr(Of(1), 2) -> 1
//	DerefOr(nil, 2)   -> 2
func DerefOr[E any](ptr *E, or E) E {
	if ptr != nil {
		return *ptr
	}
	return or
}

// New returns a non-nil pointer to zero value of type E.
// Function is useful as instantiation of typed 'new function.
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
