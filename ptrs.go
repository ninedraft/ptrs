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

// Flatten returns values, obtained by dereferencing provided pointers.
// Nil pointers will be skipped.
// Order of elements will be preseverved.
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

// Ref takes pointers of provided slice elements.
// Order of elements is preserved.
func Ref[E any](values []E) []*E {
	var ptrs = make([]*E, len(values))
	for i := range values {
		ptrs[i] = &values[i]
	}
	return ptrs
}

// Make allocates slice of n elements, takes pointer of each one, 
// and calls fn(i, ptr), where i is element index and ptr=&elements[i]. 
// If fn is nil, then just returns pointers.
//
// Make function makes a big one allocation instead of creating elements one-by-one in a cycle,
// which is much GC friendly approach and performs nearly a 2x faster.
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

// EqualSlice compares two slices of pointers.
// Slices are equal if have equal lengths and each pair of pointers
// are equal or pointing to equal values.
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

// Equal returns true, if a==b or *a==*b.
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
